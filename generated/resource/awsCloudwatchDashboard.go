package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudwatchDashboard = `{
  "block": {
    "attributes": {
      "dashboard_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dashboard_body": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dashboard_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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

func AwsCloudwatchDashboardSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudwatchDashboard), &result)
	return &result
}
