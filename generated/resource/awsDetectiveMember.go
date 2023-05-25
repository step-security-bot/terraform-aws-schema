package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDetectiveMember = `{
  "block": {
    "attributes": {
      "account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "administrator_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disable_email_notification": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disabled_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "email_address": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "graph_arn": {
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
      "invited_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "message": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "updated_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "volume_usage_in_bytes": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDetectiveMemberSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDetectiveMember), &result)
	return &result
}
