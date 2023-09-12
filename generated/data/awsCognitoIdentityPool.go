package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCognitoIdentityPool = `{
  "block": {
    "attributes": {
      "allow_classic_flow": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "allow_unauthenticated_identities": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cognito_identity_providers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "client_id": "string",
              "provider_name": "string",
              "server_side_token_check": "bool"
            }
          ]
        ]
      },
      "developer_provider_name": {
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
      "identity_pool_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "openid_connect_provider_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "saml_provider_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "supported_login_providers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
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

func AwsCognitoIdentityPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCognitoIdentityPool), &result)
	return &result
}
