package motor

import (
	"fmt"

	ct "github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/sonr-io/sonr/x/schema/types"
	_ "golang.org/x/mobile/bind"
)

func QuerySchema(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.QueryWhatIsRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	res, err := instance.GetClient().QueryWhatIs(&types.QueryWhatIsRequest{
		Did:     request.Did,
		Creator: request.Creator,
	})
	if err != nil {
		return nil, err
	}
	return res.Marshal()
}

func QuerySchemaByCreator(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.QueryWhatIsByCreatorRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	res, err := instance.GetClient().QueryWhatIsByCreator(&types.QueryWhatIsCreatorRequest{
		Creator: request.Creator,
	})
	if err != nil {
		return nil, err
	}
	return res.Marshal()
}

func QuerySchemaByDid(did string) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	res, err := instance.GetClient().QueryWhatIsByDid(&types.QueryWhatIsByDidRequest{
		Did: did,
	})
	if err != nil {
		return nil, err
	}
	return res.Marshal()
}
