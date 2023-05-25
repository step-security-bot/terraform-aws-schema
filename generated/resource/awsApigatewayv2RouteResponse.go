package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApigatewayv2RouteResponse = `{
  "block": {
    "attributes": {
      "api_id": {
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
      "model_selection_expression": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "response_models": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "route_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "route_response_key": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsApigatewayv2RouteResponseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApigatewayv2RouteResponse), &result)
	return &result
}
