package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSyntheticsCanary = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "artifact_s3_location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "engine_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "execution_role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "failure_retention_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "handler": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "runtime_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_bucket": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_location_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "start_canary": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "success_retention_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      },
      "timeline": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "created": "string",
              "last_modified": "string",
              "last_started": "string",
              "last_stopped": "string"
            }
          ]
        ]
      },
      "zip_file": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "artifact_config": {
        "block": {
          "block_types": {
            "s3_encryption": {
              "block": {
                "attributes": {
                  "encryption_mode": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key_arn": {
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "run_config": {
        "block": {
          "attributes": {
            "active_tracing": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "memory_in_mb": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "timeout_in_seconds": {
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
      "schedule": {
        "block": {
          "attributes": {
            "duration_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "expression": {
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
      },
      "vpc_config": {
        "block": {
          "attributes": {
            "security_group_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnet_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "vpc_id": {
              "computed": true,
              "description_kind": "plain",
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

func AwsSyntheticsCanarySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSyntheticsCanary), &result)
	return &result
}
