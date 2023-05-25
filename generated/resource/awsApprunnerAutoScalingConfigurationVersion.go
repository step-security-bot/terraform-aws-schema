package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApprunnerAutoScalingConfigurationVersion = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_scaling_configuration_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "auto_scaling_configuration_revision": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "latest": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "max_concurrency": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsApprunnerAutoScalingConfigurationVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApprunnerAutoScalingConfigurationVersion), &result)
	return &result
}
