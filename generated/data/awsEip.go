package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEip = `{
  "block": {
    "attributes": {
      "association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "carrier_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_owned_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_owned_ipv4_pool": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain": {
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
      "instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_interface_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_interface_owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_dns": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_dns": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_ip": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_ipv4_pool": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEipSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEip), &result)
	return &result
}
