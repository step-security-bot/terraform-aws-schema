package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-aws-schema/v4/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsEc2TransitGatewayRouteTablesSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsEc2TransitGatewayRouteTablesSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}