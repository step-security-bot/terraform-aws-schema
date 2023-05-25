package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBackupPlan = `{
  "block": {
    "attributes": {
      "arn": {
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
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "rule": {
        "block": {
          "attributes": {
            "completion_window": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "recovery_point_tags": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "rule_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "schedule": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start_window": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "target_vault_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "copy_action": {
              "block": {
                "attributes": {
                  "destination_vault_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "lifecycle": {
                    "block": {
                      "attributes": {
                        "cold_storage_after": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "delete_after": {
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
              "nesting_mode": "set"
            },
            "lifecycle": {
              "block": {
                "attributes": {
                  "cold_storage_after": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "delete_after": {
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
        "min_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBackupPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBackupPlan), &result)
	return &result
}
