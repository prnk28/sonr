package motor

import (
	"fmt"
	"log"

	st "github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/assert"
)

func (suite *MotorTestSuite) Test_CreateSchema() {
	fmt.Println("Test_CreateSchema")
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

	resp, err := suite.motorWithKeys.GetClient().CreateSchema(createSchemaRequest)
	assert.NoError(suite.T(), err, "schema created successfully")
	fmt.Printf("success: %s\n", resp.WhatIs)
}

func (suite *MotorTestSuite) Test_QuerySchema() {
	fmt.Println("Test_QuerySchema")
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

	resp, err := suite.motorWithKeys.GetClient().CreateSchema(createSchemaRequest)
	assert.NoError(suite.T(), err, "schema created successfully")

	// CREATE DONE, TRY QUERY
	qReq := &st.QueryWhatIsRequest{
		Creator: suite.motorWithKeys.GetAddress(),
		Did:     resp.WhatIs.Did,
	}

	qresp, err := suite.motorWithKeys.GetClient().QueryWhatIs(qReq)
	assert.NoError(suite.T(), err, "query response succeeds")
	assert.Equal(suite.T(), resp.WhatIs.Did, qresp.WhatIs.Did)
}

func (suite *MotorTestSuite) Test_QuerySchemaByCreator() {
	fmt.Println("Test_QuerySchemaByCreator")
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

	resp, err := suite.motorWithKeys.GetClient().CreateSchema(createSchemaRequest)
	assert.NoError(suite.T(), err, "schema created successfully")

	qReq := &st.QueryWhatIsCreatorRequest{
		Creator: resp.WhatIs.Creator,
	}

	qresp, err := suite.motorWithKeys.GetClient().QueryWhatIsByCreator(qReq)
	assert.NoError(suite.T(), err, "query response succeeds")
	if err != nil {
		log.Fatal(err)
	}

	if qresp.WhatIs != nil {
		fmt.Println(qresp.WhatIs)
	} else {
		fmt.Println("no schemas.")
	}
}

func (suite *MotorTestSuite) Test_QuerySchemaByDid() {
	fmt.Println("Test_QuerySchemaByDid")
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

	resp, err := suite.motorWithKeys.GetClient().CreateSchema(createSchemaRequest)
	assert.NoError(suite.T(), err, "schema created successfully")

	// CREATE DONE, TRY QUERY
	qresp, err := suite.motorWithKeys.GetClient().QueryWhatIsByDid(&st.QueryWhatIsByDidRequest{Did: resp.WhatIs.Did})
	assert.NoError(suite.T(), err, "query response succeeds")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(qresp)
}
