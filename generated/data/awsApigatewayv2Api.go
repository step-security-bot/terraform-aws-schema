package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApigatewayv2Api = `{
  "block": {
    "attributes": {
      "api_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "api_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "api_key_selection_expression": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cors_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allow_credentials": "bool",
              "allow_headers": [
                "set",
                "string"
              ],
              "allow_methods": [
                "set",
                "string"
              ],
              "allow_origins": [
                "set",
                "string"
              ],
              "expose_headers": [
                "set",
                "string"
              ],
              "max_age": "number"
            }
          ]
        ]
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disable_execute_api_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "execution_arn": {
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "protocol_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "route_selection_expression": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsApigatewayv2ApiSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApigatewayv2Api), &result)
	return &result
}
