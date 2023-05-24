package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpcNetworkPerformanceMetricSubscription = `{
  "block": {
    "attributes": {
      "destination": {
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
      "metric": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "period": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "statistic": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVpcNetworkPerformanceMetricSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpcNetworkPerformanceMetricSubscription), &result)
	return &result
}
