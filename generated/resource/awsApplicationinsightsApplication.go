package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApplicationinsightsApplication = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_config_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_create": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cwe_monitor_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "grouping_type": {
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
      "ops_center_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ops_item_sns_topic_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
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

func AwsApplicationinsightsApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApplicationinsightsApplication), &result)
	return &result
}
