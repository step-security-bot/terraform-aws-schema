package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlobalacceleratorAccelerator = `{
  "block": {
    "attributes": {
      "dns_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "hosted_zone_id": {
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
      "ip_address_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_sets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ip_addresses": [
                "list",
                "string"
              ],
              "ip_family": "string"
            }
          ]
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "attributes": {
        "block": {
          "attributes": {
            "flow_logs_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "flow_logs_s3_bucket": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "flow_logs_s3_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGlobalacceleratorAcceleratorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlobalacceleratorAccelerator), &result)
	return &result
}
