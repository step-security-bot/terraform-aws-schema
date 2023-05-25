package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSnsTopicSubscription = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "confirmation_timeout_in_minutes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "delivery_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_auto_confirms": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "filter_policy": {
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
      "protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "raw_message_delivery": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "topic_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSnsTopicSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSnsTopicSubscription), &result)
	return &result
}
