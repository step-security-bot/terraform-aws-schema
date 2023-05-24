package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNeptuneEngineVersion = `{
  "block": {
    "attributes": {
      "engine": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "exportable_log_types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameter_group_family": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_versions": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "supported_timezones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "supports_log_exports_to_cloudwatch": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_read_replica": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "valid_upgrade_targets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version_description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNeptuneEngineVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNeptuneEngineVersion), &result)
	return &result
}
