package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApigatewayv2Integration = `{
  "block": {
    "attributes": {
      "api_id": {
        "description_kind": "plain",
        "required": true,
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
      "content_handling_strategy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "credentials_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
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
      "integration_method": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "integration_response_selection_expression": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "integration_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "integration_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "passthrough_behavior": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payload_format_version": {
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
      "template_selection_expression": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "timeout_milliseconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsApigatewayv2IntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApigatewayv2Integration), &result)
	return &result
}
