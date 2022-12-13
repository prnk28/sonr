package mpc

import (
	"errors"
	"sync"

	"github.com/sonr-io/multi-party-sig/pkg/ecdsa"
	"github.com/sonr-io/multi-party-sig/pkg/math/curve"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/pkg/pool"
	"github.com/sonr-io/multi-party-sig/pkg/protocol"
	"github.com/sonr-io/multi-party-sig/protocols/cmp"
	"github.com/sonr-io/sonr/internal/node"
)

func cmpKeygen(id party.ID, ids node.IDSlice, n *node.Channel, threshold int, wg *sync.WaitGroup, pl *pool.Pool) (*cmp.Config, error) {
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Keygen(curve.Secp256k1{}, id, ids.ToPartyIDSlice(), threshold, pl), nil)
	if err != nil {
		return nil, err
	}

	handlerLoopChannel(id, h, n)
	r, err := h.Result()
	if err != nil {
		return nil, err
	}
	conf := r.(*cmp.Config)
	return conf, nil
}

<<<<<<<< HEAD:core/vault/x/mpc/cmp.go
func cmpRefresh(c *cmp.Config, n *node.Channel, wg *sync.WaitGroup, pl *pool.Pool) (*cmp.Config, error) {
========
func cmpRefresh(c *cmp.Config, n *Network, wg *sync.WaitGroup, pl *pool.Pool) (*cmp.Config, error) {
>>>>>>>> master:pkg/protocol/vault/x/mpc/cmp.go
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Refresh(c, pl), nil)
	if err != nil {
		return nil, err
	}
<<<<<<<< HEAD:core/vault/x/mpc/cmp.go
	handlerLoopChannel(c.ID, h, n)
========
	handlerLoop(c.ID, h, n)
>>>>>>>> master:pkg/protocol/vault/x/mpc/cmp.go
	r, err := h.Result()
	if err != nil {
		return nil, err
	}
	conf := r.(*cmp.Config)
	return conf, nil
}

<<<<<<<< HEAD:core/vault/x/mpc/cmp.go
func cmpSign(c *cmp.Config, m []byte, signers party.IDSlice, n *node.Channel, wg *sync.WaitGroup, pl *pool.Pool) (*ecdsa.Signature, error) {
========
func cmpSign(c *cmp.Config, m []byte, signers party.IDSlice, n *Network, wg *sync.WaitGroup, pl *pool.Pool) (*ecdsa.Signature, error) {
>>>>>>>> master:pkg/protocol/vault/x/mpc/cmp.go
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Sign(c, signers, m, pl), nil)
	if err != nil {
		return nil, err
	}
	handlerLoopChannel(c.ID, h, n)

	signResult, err := h.Result()
	if err != nil {
		return nil, err
	}
	signature := signResult.(*ecdsa.Signature)
	if !signature.Verify(c.PublicPoint(), m) {
		return nil, errors.New("failed to verify cmp signature")
	}
	return signature, nil
}

<<<<<<<< HEAD:core/vault/x/mpc/cmp.go
func cmpPreSign(c *cmp.Config, signers party.IDSlice, n *node.Channel, wg *sync.WaitGroup, pl *pool.Pool) (*ecdsa.PreSignature, error) {
========
func cmpPreSign(c *cmp.Config, signers party.IDSlice, n *Network, wg *sync.WaitGroup, pl *pool.Pool) (*ecdsa.PreSignature, error) {
>>>>>>>> master:pkg/protocol/vault/x/mpc/cmp.go
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Presign(c, signers, pl), nil)
	if err != nil {
		return nil, err
	}
	handlerLoopChannel(c.ID, h, n)
	signResult, err := h.Result()
	if err != nil {
		return nil, err
	}
	preSignature := signResult.(*ecdsa.PreSignature)
	if err = preSignature.Validate(); err != nil {
		return nil, errors.New("failed to verify cmp presignature")
	}
	return preSignature, nil
}

<<<<<<<< HEAD:core/vault/x/mpc/cmp.go
func cmpPreSignOnline(c *cmp.Config, preSignature *ecdsa.PreSignature, m []byte, n *node.Channel, wg *sync.WaitGroup, pl *pool.Pool) (*ecdsa.Signature, error) {
========
func cmpPreSignOnline(c *cmp.Config, preSignature *ecdsa.PreSignature, m []byte, n *Network, wg *sync.WaitGroup, pl *pool.Pool) (*ecdsa.Signature, error) {
>>>>>>>> master:pkg/protocol/vault/x/mpc/cmp.go
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.PresignOnline(c, preSignature, m, pl), nil)
	if err != nil {
		return nil, err
	}
	handlerLoopChannel(c.ID, h, n)

	signResult, err := h.Result()
	if err != nil {
		return nil, err
	}
	signature := signResult.(*ecdsa.Signature)
	if !signature.Verify(c.PublicPoint(), m) {
		return nil, errors.New("failed to verify cmp signature")
	}
	return signature, nil
}
