package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIotAuthorizer = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_function_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enable_caching_for_http": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "signing_disabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "token_key_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "token_signing_public_keys": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
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

func AwsIotAuthorizerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIotAuthorizer), &result)
	return &result
}
