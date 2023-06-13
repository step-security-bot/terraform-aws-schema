package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRoute53TrafficPolicyInstance = `{
  "block": {
    "attributes": {
      "hosted_zone_id": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "traffic_policy_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "traffic_policy_version": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "ttl": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRoute53TrafficPolicyInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRoute53TrafficPolicyInstance), &result)
	return &result
}
