package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLakeformationDataLakeSettings = `{
  "block": {
    "attributes": {
      "admins": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "catalog_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_database_default_permissions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "permissions": [
                "set",
                "string"
              ],
              "principal": "string"
            }
          ]
        ]
      },
      "create_table_default_permissions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "permissions": [
                "set",
                "string"
              ],
              "principal": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "trusted_resource_owners": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLakeformationDataLakeSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLakeformationDataLakeSettings), &result)
	return &result
}
