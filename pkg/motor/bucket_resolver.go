package motor

import (
	"context"

	"github.com/sonr-io/sonr/internal/bucket"
	mt "github.com/sonr-io/sonr/third_party/types/motor"
)

func (mtr *motorNodeImpl) GetBucket(did string) (bucket.Bucket, error) {
	addr := mtr.GetAddress()
	if _, ok := mtr.Resources.whereIsStore[did]; !ok {
		qreq := mt.QueryWhereIsRequest{
			Creator: addr,
			Did:     did,
		}
		_, err := mtr.QueryWhereIs(qreq)
		wi := mtr.Resources.whereIsStore[did]

		if err != nil {
			return nil, err
		}
		b := bucket.New(addr, wi, mtr.Resources.shell, mtr.GetClient())

		err = b.ResolveBuckets()
		if err != nil {
			return nil, err
		}
		err = b.ResolveContent()
		if err != nil {
			return nil, err
		}

		mtr.Resources.bucketStore[did] = b
		for _, sb := range b.GetBuckets() {
			mtr.Resources.bucketStore[sb.GetDID()] = sb
		}
	}

	return mtr.Resources.bucketStore[did], nil
}

/*
Takes the whereIs store and checks for a matching bucket in the cache, if its not present it will create it and get its sub buckets
Does not query for new buckets, only respects what is currently present in the store
*/
func (mtr *motorNodeImpl) GetBuckets(context context.Context) ([]bucket.Bucket, error) {
	addr := mtr.GetAddress()

	var buckets []bucket.Bucket = make([]bucket.Bucket, len(mtr.Resources.whereIsStore))
	for _, wi := range mtr.Resources.whereIsStore {
		did := wi.Did
		if _, ok := mtr.Resources.bucketStore[did]; !ok {
			b := bucket.New(addr, wi, mtr.Resources.shell, mtr.GetClient())

			err := b.ResolveBuckets()
			if err != nil {
				return nil, err
			}
			err = b.ResolveContent()
			if err != nil {
				return nil, err
			}

			mtr.Resources.bucketStore[did] = b
			for _, sb := range b.GetBuckets() {
				mtr.Resources.bucketStore[sb.GetDID()] = sb
			}
		}
		buckets = append(buckets, mtr.Resources.bucketStore[did])
	}

	return buckets, nil
}