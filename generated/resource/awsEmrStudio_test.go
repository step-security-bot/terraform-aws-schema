package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v4/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsEmrStudioSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsEmrStudioSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}