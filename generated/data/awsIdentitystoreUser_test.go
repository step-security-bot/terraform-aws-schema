package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v4/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsIdentitystoreUserSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsIdentitystoreUserSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}