package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGameliftFleet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "build_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "build_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ec2_instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "fleet_type": {
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
      "instance_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_paths": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "metric_groups": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "new_game_session_protection_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "operating_system": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "script_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "script_id": {
        "description_kind": "plain",
        "optional": true,
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
      "certificate_configuration": {
        "block": {
          "attributes": {
            "certificate_type": {
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
      "ec2_inbound_permission": {
        "block": {
          "attributes": {
            "from_port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "ip_range": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "to_port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 50,
        "nesting_mode": "set"
      },
      "resource_creation_limit_policy": {
        "block": {
          "attributes": {
            "new_game_sessions_per_creator": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "policy_period_in_minutes": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "runtime_configuration": {
        "block": {
          "attributes": {
            "game_session_activation_timeout_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_concurrent_game_session_activations": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "server_process": {
              "block": {
                "attributes": {
                  "concurrent_executions": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "launch_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "parameters": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 50,
              "nesting_mode": "list"
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

func AwsGameliftFleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGameliftFleet), &result)
	return &result
}
