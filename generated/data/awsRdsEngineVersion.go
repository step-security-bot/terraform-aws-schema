package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRdsEngineVersion = `{
  "block": {
    "attributes": {
      "default_character_set": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine": {
        "description_kind": "plain",
        "required": true,
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
          "set",
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
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "supported_character_sets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "supported_feature_names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "supported_modes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
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
      "supports_global_databases": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_log_exports_to_cloudwatch": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_parallel_query": {
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

func AwsRdsEngineVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRdsEngineVersion), &result)
	return &result
}
