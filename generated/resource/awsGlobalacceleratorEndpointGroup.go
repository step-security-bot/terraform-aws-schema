package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlobalacceleratorEndpointGroup = `{
  "block": {
    "attributes": {
      "endpoint_group_region": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_interval_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_check_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "health_check_protocol": {
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
      "listener_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "threshold_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "traffic_dial_percentage": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "endpoint_configuration": {
        "block": {
          "attributes": {
            "endpoint_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "weight": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 10,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGlobalacceleratorEndpointGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlobalacceleratorEndpointGroup), &result)
	return &result
}
