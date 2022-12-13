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

// func cmpRefresh(c *cmp.Config, n *node.Channel, wg *sync.WaitGroup, pl *pool.Pool) (*cmp.Config, error) {

// 	handlerLoopChannel(c.ID, h, n)
// 	r, err := h.Result()
// 	if err != nil {
// 		return nil, err
// 	}
// 	conf := r.(*cmp.Config)
// 	return conf, nil
// }

func cmpSign(c *cmp.Config, m []byte, signers party.IDSlice, n *node.Channel, wg *sync.WaitGroup, pl *pool.Pool) (*ecdsa.Signature, error) {
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

func cmpPreSign(c *cmp.Config, signers party.IDSlice, n *node.Channel, wg *sync.WaitGroup, pl *pool.Pool) (*ecdsa.PreSignature, error) {
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

func cmpPreSignOnline(c *cmp.Config, preSignature *ecdsa.PreSignature, m []byte, n *node.Channel, wg *sync.WaitGroup, pl *pool.Pool) (*ecdsa.Signature, error) {
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
