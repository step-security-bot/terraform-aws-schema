package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsKmsSecrets = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "plaintext": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "secret": {
        "block": {
          "attributes": {
            "context": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "grant_tokens": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "payload": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsKmsSecretsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsKmsSecrets), &result)
	return &result
}
