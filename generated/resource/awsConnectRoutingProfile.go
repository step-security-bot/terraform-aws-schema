package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsConnectRoutingProfile = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_outbound_queue_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "routing_profile_id": {
        "computed": true,
        "description_kind": "plain",
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
      "media_concurrencies": {
        "block": {
          "attributes": {
            "channel": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "concurrency": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      },
      "queue_configs": {
        "block": {
          "attributes": {
            "channel": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "delay": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "priority": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "queue_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "queue_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "queue_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsConnectRoutingProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsConnectRoutingProfile), &result)
	return &result
}
