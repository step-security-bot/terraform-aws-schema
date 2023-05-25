package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsPinpointApp = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "campaign_hook": {
        "block": {
          "attributes": {
            "lambda_function_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "web_url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "limits": {
        "block": {
          "attributes": {
            "daily": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "maximum_duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "messages_per_second": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "total": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "quiet_time": {
        "block": {
          "attributes": {
            "end": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsPinpointAppSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsPinpointApp), &result)
	return &result
}
