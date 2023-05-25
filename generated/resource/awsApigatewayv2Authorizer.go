package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApigatewayv2Authorizer = `{
  "block": {
    "attributes": {
      "api_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "authorizer_credentials_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "authorizer_payload_format_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "authorizer_result_ttl_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "authorizer_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "authorizer_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_simple_responses": {
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
      "identity_sources": {
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
      }
    },
    "block_types": {
      "jwt_configuration": {
        "block": {
          "attributes": {
            "audience": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "issuer": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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

func AwsApigatewayv2AuthorizerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApigatewayv2Authorizer), &result)
	return &result
}
