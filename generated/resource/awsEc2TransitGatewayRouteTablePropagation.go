package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2TransitGatewayRouteTablePropagation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_attachment_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_gateway_route_table_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEc2TransitGatewayRouteTablePropagationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2TransitGatewayRouteTablePropagation), &result)
	return &result
}
