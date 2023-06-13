package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v5/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsS3BucketObjectsSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsS3BucketObjectsSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
