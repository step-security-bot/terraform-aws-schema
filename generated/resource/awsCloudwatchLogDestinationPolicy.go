package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudwatchLogDestinationPolicy = `{
  "block": {
    "attributes": {
      "access_policy": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "force_update": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudwatchLogDestinationPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudwatchLogDestinationPolicy), &result)
	return &result
}
