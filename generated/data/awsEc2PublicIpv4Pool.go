package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEc2PublicIpv4Pool = `{
  "block": {
    "attributes": {
      "description": {
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
      "network_border_group": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pool_address_ranges": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address_count": "number",
              "available_address_count": "number",
              "first_address": "string",
              "last_address": "string"
            }
          ]
        ]
      },
      "pool_id": {
        "description_kind": "plain",
        "required": true,
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
      },
      "total_address_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "total_available_address_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEc2PublicIpv4PoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEc2PublicIpv4Pool), &result)
	return &result
}
