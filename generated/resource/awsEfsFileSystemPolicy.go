package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEfsFileSystemPolicy = `{
  "block": {
    "attributes": {
      "bypass_policy_lockout_safety_check": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "file_system_id": {
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
      "policy": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsEfsFileSystemPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEfsFileSystemPolicy), &result)
	return &result
}
