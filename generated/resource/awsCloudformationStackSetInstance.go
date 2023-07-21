package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudformationStackSetInstance = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "call_as": {
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
      "organizational_unit_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "parameter_overrides": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "region": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retain_stack": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "stack_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stack_instance_summaries": {
        "computed": true,
        "description": "List of stack instances created from an organizational unit deployment target. This will only be populated when ` + "`" + `deployment_targets` + "`" + ` is set.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "account_id": "string",
              "organizational_unit_id": "string",
              "stack_id": "string"
            }
          ]
        ]
      },
      "stack_set_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "deployment_targets": {
        "block": {
          "attributes": {
            "organizational_unit_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "operation_preferences": {
        "block": {
          "attributes": {
            "failure_tolerance_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "failure_tolerance_percentage": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_concurrent_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_concurrent_percentage": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "region_concurrency_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region_order": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudformationStackSetInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudformationStackSetInstance), &result)
	return &result
}
