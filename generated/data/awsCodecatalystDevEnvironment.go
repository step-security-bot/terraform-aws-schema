package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCodecatalystDevEnvironment = `{
  "block": {
    "attributes": {
      "alias": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creator_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "env_id": {
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
      "ides": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string",
              "runtime": "string"
            }
          ]
        ]
      },
      "inactivity_timeout_minutes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "persistent_storage": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "size": "number"
            }
          ]
        ]
      },
      "project_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "space_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
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
      "repositories": {
        "block": {
          "attributes": {
            "branch_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "repository_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 100,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCodecatalystDevEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCodecatalystDevEnvironment), &result)
	return &result
}
