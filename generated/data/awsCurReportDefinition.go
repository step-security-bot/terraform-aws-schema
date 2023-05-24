package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCurReportDefinition = `{
  "block": {
    "attributes": {
      "additional_artifacts": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "additional_schema_elements": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "compression": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "format": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "refresh_closed_reports": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "report_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "report_versioning": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "s3_bucket": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "s3_prefix": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "s3_region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "time_unit": {
        "computed": true,
        "description_kind": "plain",
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
