package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEmrManagedScalingPolicy = `{
  "block": {
    "attributes": {
      "cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "compute_limits": {
        "block": {
          "attributes": {
            "maximum_capacity_units": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "maximum_core_capacity_units": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "maximum_ondemand_capacity_units": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "minimum_capacity_units": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "unit_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEmrManagedScalingPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEmrManagedScalingPolicy), &result)
	return &result
}
