package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsmcontactsContactChannel = `{
  "block": {
    "attributes": {
      "activation_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "contact_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "delivery_address": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "simple_address": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSsmcontactsContactChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsmcontactsContactChannel), &result)
	return &result
}
