package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSnsPlatformApplication = `{
  "block": {
    "attributes": {
      "apple_platform_bundle_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "apple_platform_team_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "event_delivery_failure_topic_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_endpoint_created_topic_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_endpoint_deleted_topic_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_endpoint_updated_topic_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failure_feedback_role_arn": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "platform": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "platform_credential": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "platform_principal": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "success_feedback_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "success_feedback_sample_rate": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSnsPlatformApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSnsPlatformApplication), &result)
	return &result
}
