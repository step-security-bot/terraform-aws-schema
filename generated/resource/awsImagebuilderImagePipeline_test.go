package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v5/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsImagebuilderImagePipelineSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsImagebuilderImagePipelineSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
