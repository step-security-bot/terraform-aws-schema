package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlobalacceleratorListener = `{
  "block": {
    "attributes": {
      "accelerator_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "client_affinity": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "port_range": {
        "block": {
          "attributes": {
            "from_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "to_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 10,
        "min_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGlobalacceleratorListenerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlobalacceleratorListener), &result)
	return &result
}
