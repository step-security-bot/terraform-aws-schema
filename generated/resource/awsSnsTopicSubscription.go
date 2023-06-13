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
      "confirmation_was_authenticated": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      "filter_policy_scope": {
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
      "owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pending_confirmation": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      "redrive_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscription_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
