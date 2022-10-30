package motor

import (
	"fmt"

	bt "github.com/cosmos/cosmos-sdk/x/bank/types"
	mtr "github.com/sonr-io/sonr/pkg/motor"
	"github.com/sonr-io/sonr/pkg/motor/x/document"
	ct "github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	rt "github.com/sonr-io/sonr/x/registry/types"
	_ "golang.org/x/mobile/bind"
)

var (
	docBuilders map[string]*document.DocumentBuilder
	instance    mtr.MotorNode
)

type MotorCallback interface {
	OnDiscover(data []byte)
	OnWalletEvent(data []byte)
}

func Init(buf []byte, cb MotorCallback) ([]byte, error) {
	// Unmarshal the request
	var req mt.InitializeRequest
	if err := req.Unmarshal(buf); err != nil {
		return nil, err
	}

	// Create Motor instance
	mtr, err := mtr.EmptyMotor(&req, cb)
	if err != nil {
		return nil, err
	}
	instance = mtr

	// init docBuilders
	docBuilders = make(map[string]*document.DocumentBuilder)

	// Return Initialization Response
	resp := mt.InitializeResponse{
		Success: true,
	}

	if req.AuthInfo != nil {
		if res, err := instance.Login(mt.LoginRequest{
			AccountId: req.AuthInfo.Did,
			Password:  req.AuthInfo.Password,
		}); err == nil {
			return res.Marshal()
		}
	}
	return resp.Marshal()
}

func CreateAccount(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}
	// decode request
	request := mt.CreateAccountRequest{}
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	if res, err := instance.CreateAccount(request); err == nil {
		return res.Marshal()
	} else {
		return nil, err
	}
}

func CreateAccountWithKeys(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}
	// decode request
	request := mt.CreateAccountWithKeysRequest{}
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	if res, err := instance.CreateAccountWithKeys(request); err == nil {
		return res.Marshal()
	} else {
		return nil, err
	}
}

func Login(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	// decode request
	var request mt.LoginRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("error unmarshalling request: %s", err)
	}

	if res, err := instance.Login(request); err == nil {
		return res.Marshal()
	} else {
		return nil, err
	}
}

func LoginWithKeys(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	// decode request
	var request mt.LoginWithKeysRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("error unmarshalling request: %s", err)
	}

	if res, err := instance.LoginWithKeys(request); err == nil {
		return res.Marshal()
	} else {
		return nil, err
	}
}

func Connect() error {
	if instance == nil {
		return ct.ErrMotorWalletNotInitialized
	}
	return instance.Connect()
}

// SendTokens sends the specified amount of tokens to the specified address.
func SendTokens(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request bt.MsgSend
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	if res, err := instance.GetClient().SendTokens(&request); err == nil {
		return res.Marshal()
	} else {
		return nil, err
	}
}

// Stat returns general information about the Motor node its wallet and accompanying Account.
func Stat() ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	doc := instance.GetDIDDocument()
	if doc == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}
	didDoc, err := rt.NewDIDDocumentFromPkg(doc)
	if err != nil {
		return nil, err
	}

	resp := mt.StatResponse{
		Address:     instance.GetAddress(),
		Balance:     int32(instance.GetBalance()),
		DidDocument: didDoc,
	}
	return resp.Marshal()
}

func BuyAlias(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request rt.MsgBuyAlias
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	resp, err := instance.GetClient().BuyAlias(&request)
	if err != nil {
		return nil, err
	}

	return resp.Marshal()
}

func SellAlias(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request rt.MsgSellAlias
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	resp, err := instance.GetClient().SellAlias(&request)
	if err != nil {
		return nil, err
	}

	return resp.Marshal()
}

func TransferAlias(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request rt.MsgTransferAlias
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	resp, err := instance.GetClient().TransferAlias(&request)
	if err != nil {
		return nil, err
	}

	return resp.Marshal()
}
