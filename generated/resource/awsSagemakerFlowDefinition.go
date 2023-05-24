package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerFlowDefinition = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "flow_definition_name": {
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
      "role_arn": {
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
      "tags_all": {
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
      "human_loop_activation_config": {
        "block": {
          "block_types": {
            "human_loop_activation_conditions_config": {
              "block": {
                "attributes": {
                  "human_loop_activation_conditions": {
                    "description_kind": "plain",
                    "required": true,
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "human_loop_config": {
        "block": {
          "attributes": {
            "human_task_ui_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "task_availability_lifetime_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "task_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "task_description": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "task_keywords": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "task_time_limit_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "task_title": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "workteam_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "public_workforce_task_price": {
              "block": {
                "block_types": {
                  "amount_in_usd": {
                    "block": {
                      "attributes": {
                        "cents": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "dollars": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "tenth_fractions_of_a_cent": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "human_loop_request_source": {
        "block": {
          "attributes": {
            "aws_managed_human_loop_request_source": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "output_config": {
        "block": {
          "attributes": {
            "kms_key_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_output_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSagemakerFlowDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerFlowDefinition), &result)
	return &result
}
