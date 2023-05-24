package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApigatewayv2Route = `{
  "block": {
    "attributes": {
      "api_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "api_key_required": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "authorization_scopes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "authorization_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "authorizer_id": {
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
      "model_selection_expression": {
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
      "route_key": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "route_response_selection_expression": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "request_parameter": {
        "block": {
          "attributes": {
            "request_parameter_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "required": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsApigatewayv2RouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApigatewayv2Route), &result)
	return &result
}
