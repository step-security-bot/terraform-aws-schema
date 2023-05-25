package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpnConnection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_gateway_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_gateway_id": {
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
      "routes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "destination_cidr_block": "string",
              "source": "string",
              "state": "string"
            }
          ]
        ]
      },
      "static_routes_only": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "transit_gateway_attachment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tunnel1_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel1_bgp_asn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel1_bgp_holdtime": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tunnel1_cgw_inside_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel1_inside_cidr": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tunnel1_preshared_key": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "tunnel1_vgw_inside_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel2_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel2_bgp_asn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel2_bgp_holdtime": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tunnel2_cgw_inside_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel2_inside_cidr": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tunnel2_preshared_key": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "tunnel2_vgw_inside_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vgw_telemetry": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "accepted_route_count": "number",
              "last_status_change": "string",
              "outside_ip_address": "string",
              "status": "string",
              "status_message": "string"
            }
          ]
        ]
      },
      "vpn_gateway_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVpnConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpnConnection), &result)
	return &result
}
