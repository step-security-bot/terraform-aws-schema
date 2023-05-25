package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpsworksUserProfile = `{
  "block": {
    "attributes": {
      "allow_self_management": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssh_public_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssh_username": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOpsworksUserProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpsworksUserProfile), &result)
	return &result
}
