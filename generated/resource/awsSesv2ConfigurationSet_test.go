package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v5/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsSesv2ConfigurationSetSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsSesv2ConfigurationSetSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
