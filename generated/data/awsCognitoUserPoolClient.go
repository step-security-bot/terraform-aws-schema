package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCognitoUserPoolClient = `{
  "block": {
    "attributes": {
      "access_token_validity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "allowed_oauth_flows": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "allowed_oauth_flows_user_pool_client": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "allowed_oauth_scopes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "analytics_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "application_arn": "string",
              "application_id": "string",
              "external_id": "string",
              "role_arn": "string",
              "user_data_shared": "bool"
            }
          ]
        ]
      },
      "callback_urls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "client_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "client_secret": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "default_redirect_uri": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_propagate_additional_user_context_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_token_revocation": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "explicit_auth_flows": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "generate_secret": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id_token_validity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "logout_urls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "prevent_user_existence_errors": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "read_attributes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "refresh_token_validity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "supported_identity_providers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "token_validity_units": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_token": "string",
              "id_token": "string",
              "refresh_token": "string"
            }
          ]
        ]
      },
      "user_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "write_attributes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
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
