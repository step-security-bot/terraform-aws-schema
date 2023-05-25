package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2TransitGatewayRoute = `{
  "block": {
    "attributes": {
      "blackhole": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "destination_cidr_block": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_gateway_attachment_id": {
        "description_kind": "plain",
        "optional": true,
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

func AwsEc2TransitGatewayRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2TransitGatewayRoute), &result)
	return &result
}
