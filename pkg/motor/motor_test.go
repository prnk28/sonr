package motor

import (
	"encoding/hex"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"

	rt "github.com/sonr-io/sonr/x/registry/types"
	"github.com/stretchr/testify/assert"
)

func (suite *MotorTestSuite) Test_DecodeTxData() {
	data := "0A91010A242F736F6E72696F2E736F6E722E72656769737472792E4D736743726561746557686F497312691267122A736E723134373071366D3476776D6537346A376D3573326364773939357A35796E6B747A726D377A35371A31122F6469643A736E723A3134373071366D3476776D6537346A376D3573326364773939357A35796E6B747A726D377A353730BC8FA197063801"

	mcr := &rt.MsgCreateWhoIsResponse{}
	err := decodeTxResponseData(data, mcr)
	assert.NoError(suite.T(), err, "decodes tx data successfully")
	assert.Equal(suite.T(), "snr1470q6m4vwme74j7m5s2cdw995z5ynktzrm7z57", mcr.WhoIs.Owner)
}

func decodeTxResponseData(resp string, v proto.Unmarshaler) error {
	data, err := hex.DecodeString(resp)
	if err != nil {
		return err
	}

	anyWrapper := new(types.Any)
	if err := proto.Unmarshal(data, anyWrapper); err != nil {
		return err
	}

	// TODO: figure out if there's a better 'cosmos' way of doing this
	// you have to unwrap the Any twice, and the first time the bytes get decoded
	// in the 'TypeUrl' field instead of 'Value' field
	any := new(types.Any)
	if err := proto.Unmarshal([]byte(anyWrapper.TypeUrl), any); err != nil {
		return err
	}

	return v.Unmarshal(any.Value)
}
