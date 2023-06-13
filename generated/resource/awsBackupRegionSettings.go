package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsBackupRegionSettings = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_type_management_preference": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "bool"
        ]
      },
      "resource_type_opt_in_preference": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          "bool"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsBackupRegionSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsBackupRegionSettings), &result)
	return &result
}
