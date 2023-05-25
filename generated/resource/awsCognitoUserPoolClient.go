package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCognitoUserPoolClient = `{
  "block": {
    "attributes": {
      "allowed_oauth_flows": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "allowed_oauth_flows_user_pool_client": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "allowed_oauth_scopes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "callback_urls": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "client_secret": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "default_redirect_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "explicit_auth_flows": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "generate_secret": {
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
      "logout_urls": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "prevent_user_existence_errors": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "read_attributes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "refresh_token_validity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "supported_identity_providers": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "user_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "write_attributes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "analytics_configuration": {
        "block": {
          "attributes": {
            "application_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "external_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "user_data_shared": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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

func AwsCognitoUserPoolClientSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCognitoUserPoolClient), &result)
	return &result
}
