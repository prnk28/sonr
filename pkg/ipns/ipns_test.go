package ipns_test

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/pkg/ipns"
	"github.com/sonr-io/sonr/pkg/motor"
	mtu "github.com/sonr-io/sonr/testutil/motor"
	"github.com/sonr-io/sonr/third_party/types/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/sonr-io/sonr/x/schema/types"
	st "github.com/sonr-io/sonr/x/schema/types"
)

type IPNSTestSuite struct {
	suite.Suite
	motorNode motor.MotorNode
	cid       string
}

func (suite *IPNSTestSuite) SetupSuite() {
	var err error

	suite.motorNode, err = motor.EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	if err != nil {
		fmt.Printf("Failed to setup test motor: %s\n", err)
	}

	err = mtu.SetupTestAddressWithKeys(suite.motorNode)
	if err != nil {
		fmt.Printf("Failed to setup test address: %s\n", err)
	}

	// create document
	createSchemaRequest := &st.MsgCreateSchema{
		Label: "TestUser",
		Fields: []*st.SchemaField{
			{
				Name: "email",
				FieldKind: &st.SchemaFieldKind{
					Kind: st.Kind_STRING,
				},
			},
			{
				Name: "firstName",
				FieldKind: &st.SchemaFieldKind{
					Kind: st.Kind_STRING,
				},
			},
			{
				Name: "age",
				FieldKind: &st.SchemaFieldKind{
					Kind: st.Kind_INT,
				},
			},
		},
	}

	resp, err := suite.motorNode.GetClient().CreateSchema(createSchemaRequest)
	if err != nil {
		fmt.Printf("Failed to create schema: %s\n", err)
	}

	// query WhatIs so it's cached
	_, err = suite.motorNode.GetClient().QueryWhatIsByDid(&types.QueryWhatIsByDidRequest{
		Did: resp.WhatIs.Did,
	})
	if err != nil {
		fmt.Printf("Failed to query whatis: %s\n", err)
	}

	// upload object
	builder, err := suite.motorNode.NewDocumentBuilder(resp.WhatIs.Did)
	if err != nil {
		fmt.Printf("Failed to create document builder: %s\n", err)
	}

	builder.SetLabel("Player 1")
	builder.Set("email", "test_email")
	builder.Set("firstName", "test_name")
	builder.Set("age", 10)

	_, err = builder.Build()
	if err != nil {
		fmt.Printf("Failed to build document: %s\n", err)
	}

	result, err := builder.Upload()
	if err != nil {
		fmt.Printf("Failed to upload document: %s\n", err)
	}
	if "Player 1" != result.Document.Label {
		fmt.Println("Failed to upload document")
	}

	suite.cid = result.GetCid()
}

func Test_IPNS_Suite(t *testing.T) {
	suite.Run(t, new(IPNSTestSuite))
}

func (suite *IPNSTestSuite) Test_IPNS() {
	shell := shell.NewLocalShell()
	suite.T().Run("Should create ipns record", func(t *testing.T) {
		time_stamp := fmt.Sprintf("%d", time.Now().Unix())

		out_path := filepath.Join(os.TempDir(), time_stamp+".txt")

		defer os.Remove(out_path)
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		assert.NoError(t, err)
		rec, err := ipns.New(priv)
		assert.NoError(t, err)
		rec.Builder.SetCid(suite.cid)
		err = rec.CreateRecord()
		assert.NoError(t, err)
		srv := rec.Builder.BuildService()
		assert.NotNil(t, srv)
		fmt.Print(srv.ID)
		id, err := ipns.Publish(shell, rec)
		assert.NoError(t, err)
		str, err := ipns.Resolve(shell, id)
		assert.NoError(t, err)
		assert.NotNil(t, str)
		err = shell.Get(str, out_path)
		assert.NoError(t, err)
		buf, err := os.ReadFile(out_path)
		assert.NoError(t, err)
		fmt.Print(string(buf))
	})
}
