package motor

import (
	"fmt"

	"github.com/sonr-io/sonr/third_party/types/common"
	bt "github.com/sonr-io/sonr/x/bucket/types"
	rt "github.com/sonr-io/sonr/x/registry/types"
)

func (mtr *motorNodeImpl) triggerWalletEvent(event common.WalletEvent) error {
	if mtr.callback == nil {
		return fmt.Errorf("error callback is nil cannot trigger")
	}
	b, err := event.Marshal()

	if err != nil {
		return fmt.Errorf("Error while marshalling wallet event: \n%s", err.Error())
	}

	mtr.callback.OnWalletEvent(b)

	return nil
}

func (mtr *motorNodeImpl) fetchBucketConfig(bucket *bt.BucketConfig, uuid, creator, name string) (*bt.BucketConfig, error) {
	if bucket != nil {
		return bucket, nil
	}

	if creator != "" && name != "" {
		resp, err := mtr.Cosmos.QueryBucketByCreator(&bt.QueryGetBucketByCreatorRequest{
			Creator: creator,
		})
		if err != nil {
			return nil, err
		}

		for _, config := range resp.Buckets {
			if config.Name == name {
				return &config, nil
			}
		}
	}

	if uuid != "" {
		resp, err := mtr.Cosmos.QueryBucket(&bt.QueryGetBucketRequest{
			Uuid: uuid,
		})
		if err != nil {
			return nil, err
		}
		return &resp.Bucket, nil
	}
	return nil, fmt.Errorf("no bucket config found")
}

func (mtr *motorNodeImpl) getRegistryDIDDocument() (*rt.DIDDocument, error) {
	docBz, err := mtr.DIDDocument.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return rt.NewDIDDocumentFromBytes(docBz)
}
