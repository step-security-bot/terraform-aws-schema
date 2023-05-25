package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsTransferAccess = `{
  "block": {
    "attributes": {
      "external_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "home_directory": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "home_directory_type": {
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
      "policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "home_directory_mappings": {
        "block": {
          "attributes": {
            "entry": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 50,
        "nesting_mode": "list"
      },
      "posix_profile": {
        "block": {
          "attributes": {
            "gid": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "secondary_gids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "number"
              ]
            },
            "uid": {
              "description_kind": "plain",
              "required": true,
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
  "version": 0
}`

func AwsTransferAccessSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsTransferAccess), &result)
	return &result
}
