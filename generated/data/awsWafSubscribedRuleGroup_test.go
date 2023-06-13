package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v5/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsWafSubscribedRuleGroupSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsWafSubscribedRuleGroupSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
