package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsElasticacheUser = `{
  "block": {
    "attributes": {
      "access_string": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine": {
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
      "no_password_required": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "passwords": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": [
          "set",
          "string"
        ]
      },
      "user_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "authentication_mode": {
        "block": {
          "attributes": {
            "password_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsElasticacheUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsElasticacheUser), &result)
	return &result
}
