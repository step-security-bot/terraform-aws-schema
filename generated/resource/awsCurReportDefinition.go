package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCurReportDefinition = `{
  "block": {
    "attributes": {
      "additional_artifacts": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "additional_schema_elements": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "compression": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "format": {
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
      "report_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_region": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "time_unit": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCurReportDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCurReportDefinition), &result)
	return &result
}
