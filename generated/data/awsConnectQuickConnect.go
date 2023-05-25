package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsConnectQuickConnect = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "quick_connect_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "phone_config": [
                "list",
                [
                  "object",
                  {
                    "phone_number": "string"
                  }
                ]
              ],
              "queue_config": [
                "list",
                [
                  "object",
                  {
                    "contact_flow_id": "string",
                    "queue_id": "string"
                  }
                ]
              ],
              "quick_connect_type": "string",
              "user_config": [
                "list",
                [
                  "object",
                  {
                    "contact_flow_id": "string",
                    "user_id": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "quick_connect_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsConnectQuickConnectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsConnectQuickConnect), &result)
	return &result
}
