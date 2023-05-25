package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppstreamDirectoryConfig = `{
  "block": {
    "attributes": {
      "created_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "directory_name": {
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
      "organizational_unit_distinguished_names": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "service_account_credentials": {
        "block": {
          "attributes": {
            "account_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "account_password": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
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

func AwsAppstreamDirectoryConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppstreamDirectoryConfig), &result)
	return &result
}
