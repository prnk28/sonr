package motor

import (
	"fmt"
	"strings"

	"github.com/sonr-io/multi-party-sig/pkg/math/curve"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/protocols/cmp"
	kr "github.com/sonr-io/sonr/internal/keyring"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/vault"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	rt "github.com/sonr-io/sonr/x/registry/types"
)

func (mtr *motorNodeImpl) Login(request mt.LoginRequest) (mt.LoginResponse, error) {
	var err error
	// get PSK
	psk, err := kr.GetPSK()
	if err != nil {
		return mt.LoginResponse{}, fmt.Errorf("get PSK: %s", err)
	}

	var dsc []byte
	if request.Password == "" {
		// get DSC
		dsc, err = kr.GetDSC()
		if err != nil {
			return mt.LoginResponse{}, fmt.Errorf("get DSC: %s", err)
		}
	}

	return mtr.LoginWithKeys(mt.LoginWithKeysRequest{
		AccountId: request.AccountId,
		Password:  request.Password,
		AesDscKey: dsc,
		AesPskKey: psk,
	})
}

// LoginWithKeys creates a motor node from a LoginRequest
func (mtr *motorNodeImpl) LoginWithKeys(request mt.LoginWithKeysRequest) (mt.LoginResponse, error) {
	if request.AccountId == "" {
		return mt.LoginResponse{}, fmt.Errorf("did must be provided")
	}

	// Create Client instance
	mtr.Cosmos = client.NewClient()

	// if the given ID is an alias, first fetch the address
	var (
		whoIs *rt.WhoIs
		err   error
	)
	if strings.HasSuffix(request.AccountId, ".snr") {
		whoIsResp, err := mtr.GetClient().QueryWhoIsByAlias(&rt.QueryWhoIsAliasRequest{
			Alias: request.AccountId,
		})
		if err != nil {
			return mt.LoginResponse{}, fmt.Errorf("query WhoIs by alias: %s", err)
		}
		whoIs = whoIsResp.WhoIs
	} else {
		whoIsResp, err := mtr.GetClient().QueryWhoIs(&rt.QueryWhoIsRequest{
			Did: request.AccountId,
		})
		if err != nil {
			return mt.LoginResponse{}, fmt.Errorf("query WhoIs: %s", err)
		}
		whoIs = whoIsResp.WhoIs
	}

	// setup encryption key
	mtr.encryptionKey = request.AesPskKey

	whoIs.DidDocument.Id = did.CreateDIDFromAccount(whoIs.Owner)
	mtr.DIDDocument, err = whoIs.DidDocument.ToPkgDoc()
	if err != nil {
		return mt.LoginResponse{}, fmt.Errorf("error getting DID Document: %s", err)
	}

	// get the pubkey
	vm := mtr.DIDDocument.GetVerificationMethods().FindByID(mtr.DIDDocument.GetID())
	if vm == nil {
		return mt.LoginResponse{}, fmt.Errorf("could not find assertion method with ID %s", mtr.DIDDocument.GetID())
	}

	// fetch vault shards
	shards, err := vault.New().GetVaultShards(whoIs.Owner)
	if err != nil {
		return mt.LoginResponse{}, fmt.Errorf("error getting vault shards: %s", err)
	}

	cnfgs, err := createWalletConfigs(mtr.DeviceID, request, shards)
	if err != nil {
		return mt.LoginResponse{}, fmt.Errorf("error creating preferred config: %s", err)
	}

	// generate wallet
	if err = initMotor(mtr,
		mpc.WithConfigs(cnfgs),
		mpc.WithBase58PubKey(vm.PublicKeyBase58)); err != nil {
		return mt.LoginResponse{}, fmt.Errorf("error generating wallet: %s", err)
	}

	// assign shards
	mtr.deviceShard = shards.IssuedShards[mtr.DeviceID]
	mtr.sharedShard = shards.PskShard
	mtr.recoveryShard = shards.RecoveryShard
	mtr.unusedShards = destructureShards(shards.ShardBank)
	return mt.LoginResponse{
		Success: true,
	}, nil
}

func createWalletConfigs(id string, req mt.LoginWithKeysRequest, shards vault.Vault) (map[party.ID]*cmp.Config, error) {
	configs := make(map[party.ID]*cmp.Config)

	// if a password is provided, prefer that over the DSC
	if req.Password != "" {
		// build recovery Config
		recShard, err := mpc.AesDecryptWithPassword(req.Password, shards.RecoveryShard)
		if err != nil {
			return nil, fmt.Errorf("error decrypting recovery shard (%s): %s", recShard, err)
		}

		configs["recovery"], err = hydrateConfig(recShard)
		if err != nil {
			return nil, fmt.Errorf("recovery shard: %s", err)
		}
	} else {
		// build DSC Config if password is not provided
		deviceShard, ok := shards.IssuedShards[id]
		if !ok {
			return nil, fmt.Errorf("could not find device shard with key '%s'", id)
		}
		dscShard, err := mpc.AesDecryptWithKey(req.AesDscKey, deviceShard)
		if err != nil {
			return nil, fmt.Errorf("error decrypting DSC shard: %s", err)
		}

		configs["dsc"], err = hydrateConfig(dscShard)
		if err != nil {
			return nil, fmt.Errorf("dsc shard: %s", err)
		}
	}

	// in all cases, use the PSK
	pskShard, err := mpc.AesDecryptWithKey(req.AesPskKey, shards.PskShard)
	if err != nil {
		return nil, fmt.Errorf("error decrypting PSK shard: %s", err)
	}

	configs["psk"], err = hydrateConfig(pskShard)
	if err != nil {
		return nil, fmt.Errorf("psk shard: %s", err)
	}

	return configs, nil
}

func hydrateConfig(c []byte) (*cmp.Config, error) {
	cnf := cmp.EmptyConfig(curve.Secp256k1{})
	if err := cnf.UnmarshalBinary(c); err != nil {
		return nil, fmt.Errorf("error unmarshalling shard: %s", err)
	}
	return cnf, nil
}

func destructureShards(s []vault.Shard) [][]byte {
	result := make([][]byte, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}
