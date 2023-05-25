package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCognitoIdentityPoolRolesAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identity_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "roles": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "role_mapping": {
        "block": {
          "attributes": {
            "ambiguous_role_resolution": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "identity_provider": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "mapping_rule": {
              "block": {
                "attributes": {
                  "claim": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "match_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 25,
              "nesting_mode": "list"
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

func AwsCognitoIdentityPoolRolesAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCognitoIdentityPoolRolesAttachment), &result)
	return &result
}
