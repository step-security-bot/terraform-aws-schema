package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApiGatewayAccount = `{
  "block": {
    "attributes": {
      "api_key_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cloudwatch_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "features": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "throttle_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "burst_limit": "number",
              "rate_limit": "number"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsApiGatewayAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApiGatewayAccount), &result)
	return &result
}
