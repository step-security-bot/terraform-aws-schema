package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRamResourceShareAccepter = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "invitation_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "receiver_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resources": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "sender_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "share_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "share_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "share_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
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

func AwsRamResourceShareAccepterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRamResourceShareAccepter), &result)
	return &result
}
