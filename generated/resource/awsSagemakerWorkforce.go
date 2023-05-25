package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSagemakerWorkforce = `{
  "block": {
    "attributes": {
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
      "subdomain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "workforce_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "cognito_config": {
        "block": {
          "attributes": {
            "client_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "user_pool": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "oidc_config": {
        "block": {
          "attributes": {
            "authorization_endpoint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "client_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "client_secret": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "issuer": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "jwks_uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "logout_endpoint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "token_endpoint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "user_info_endpoint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "source_ip_config": {
        "block": {
          "attributes": {
            "cidrs": {
              "description_kind": "plain",
              "required": true,
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSagemakerWorkforceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSagemakerWorkforce), &result)
	return &result
}
