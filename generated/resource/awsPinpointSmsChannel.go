package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsPinpointSmsChannel = `{
  "block": {
    "attributes": {
      "application_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "promotional_messages_per_second": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "sender_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "short_code": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transactional_messages_per_second": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsPinpointSmsChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsPinpointSmsChannel), &result)
	return &result
}
