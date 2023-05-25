package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApiGatewayIntegration = `{
  "block": {
    "attributes": {
      "cache_key_parameters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "cache_namespace": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connection_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connection_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_handling": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "credentials": {
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
      "integration_http_method": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "passthrough_behavior": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "request_parameters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "request_parameters_in_json": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "request_templates": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
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
      },
      "timeout_milliseconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsApiGatewayIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApiGatewayIntegration), &result)
	return &result
}
