package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsChimeVoiceConnectorOrigination = `{
  "block": {
    "attributes": {
      "disabled": {
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
      "voice_connector_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "route": {
        "block": {
          "attributes": {
            "host": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "priority": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "weight": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 20,
        "min_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsChimeVoiceConnectorOriginationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsChimeVoiceConnectorOrigination), &result)
	return &result
}
