package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDirectoryServiceRadiusSettings = `{
  "block": {
    "attributes": {
      "authentication_protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "directory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_label": {
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
      "radius_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "radius_retries": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "radius_servers": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "radius_timeout": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "shared_secret": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "use_same_username": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
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

func AwsDirectoryServiceRadiusSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDirectoryServiceRadiusSettings), &result)
	return &result
}
