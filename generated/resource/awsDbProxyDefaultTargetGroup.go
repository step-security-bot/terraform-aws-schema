package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDbProxyDefaultTargetGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_proxy_name": {
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "connection_pool_config": {
        "block": {
          "attributes": {
            "connection_borrow_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "init_query": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_connections_percent": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_idle_connections_percent": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "session_pinning_filters": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDbProxyDefaultTargetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDbProxyDefaultTargetGroup), &result)
	return &result
}
