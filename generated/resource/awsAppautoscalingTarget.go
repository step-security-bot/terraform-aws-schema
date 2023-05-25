package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppautoscalingTarget = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_capacity": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "min_capacity": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "resource_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scalable_dimension": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_namespace": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAppautoscalingTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppautoscalingTarget), &result)
	return &result
}
