package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v4/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsWafregionalRateBasedRuleSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsWafregionalRateBasedRuleSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}