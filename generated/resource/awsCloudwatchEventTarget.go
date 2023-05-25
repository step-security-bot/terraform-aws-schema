package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudwatchEventTarget = `{
  "block": {
    "attributes": {
      "arn": {
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
      "input": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "input_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "batch_target": {
        "block": {
          "attributes": {
            "array_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "job_attempts": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "job_definition": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "job_name": {
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
      "ecs_target": {
        "block": {
          "attributes": {
            "group": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "launch_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "platform_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "task_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "task_definition_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "network_configuration": {
              "block": {
                "attributes": {
                  "assign_public_ip": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "security_groups": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "subnets": {
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
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "input_transformer": {
        "block": {
          "attributes": {
            "input_paths": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "input_template": {
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
      "kinesis_target": {
        "block": {
          "attributes": {
            "partition_key_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "run_command_targets": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 5,
        "nesting_mode": "list"
      },
      "sqs_target": {
        "block": {
          "attributes": {
            "message_group_id": {
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

func AwsCloudwatchEventTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudwatchEventTarget), &result)
	return &result
}
