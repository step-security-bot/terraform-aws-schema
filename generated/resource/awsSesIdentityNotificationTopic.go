package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSesIdentityNotificationTopic = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identity": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "include_original_headers": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "notification_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "topic_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSesIdentityNotificationTopicSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSesIdentityNotificationTopic), &result)
	return &result
}
