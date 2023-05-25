package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIamUserSshKey = `{
  "block": {
    "attributes": {
      "encoding": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "fingerprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ssh_public_key_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "username": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsIamUserSshKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIamUserSshKey), &result)
	return &result
}
