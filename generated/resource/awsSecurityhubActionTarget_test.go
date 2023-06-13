package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v5/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsSecurityhubActionTargetSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsSecurityhubActionTargetSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
