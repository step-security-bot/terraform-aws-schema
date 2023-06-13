package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsConfigRemediationConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "automatic": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "config_rule_name": {
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
      "maximum_automatic_attempts": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "resource_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retry_attempt_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "target_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "execution_controls": {
        "block": {
          "block_types": {
            "ssm_controls": {
              "block": {
                "attributes": {
                  "concurrent_execution_rate_percentage": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "error_percentage": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "parameter": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "static_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "static_values": {
              "computed": true,
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
        "max_items": 25,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsConfigRemediationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsConfigRemediationConfiguration), &result)
	return &result
}
