package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2TransitGatewayVpcAttachmentAccepter = `{
  "block": {
    "attributes": {
      "appliance_mode_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_support": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "transit_gateway_attachment_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_gateway_default_route_table_association": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "transit_gateway_default_route_table_propagation": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "transit_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEc2TransitGatewayVpcAttachmentAccepterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2TransitGatewayVpcAttachmentAccepter), &result)
	return &result
}
