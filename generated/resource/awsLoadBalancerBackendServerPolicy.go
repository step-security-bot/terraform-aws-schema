package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLoadBalancerBackendServerPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "load_balancer_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_names": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLoadBalancerBackendServerPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLoadBalancerBackendServerPolicy), &result)
	return &result
}
