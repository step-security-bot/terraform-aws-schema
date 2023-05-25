package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSecurityhubStandardsControl = `{
  "block": {
    "attributes": {
      "control_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "control_status": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "control_status_updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disabled_reason": {
        "computed": true,
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
      "related_requirements": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "remediation_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "severity_rating": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "standards_control_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "title": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSecurityhubStandardsControlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSecurityhubStandardsControl), &result)
	return &result
}
