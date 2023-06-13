package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlueSchema = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compatibility": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_format": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
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
      "latest_schema_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "next_schema_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "registry_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "registry_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "schema_checkpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "schema_definition": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schema_name": {
        "description_kind": "plain",
        "required": true,
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
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGlueSchemaSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlueSchema), &result)
	return &result
}
