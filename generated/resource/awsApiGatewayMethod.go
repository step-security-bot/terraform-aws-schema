package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApiGatewayMethod = `{
  "block": {
    "attributes": {
      "api_key_required": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "authorization": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "authorization_scopes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "authorizer_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "http_method": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "operation_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "request_models": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "request_parameters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "bool"
        ]
      },
      "request_validator_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rest_api_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsApiGatewayMethodSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApiGatewayMethod), &result)
	return &result
}
