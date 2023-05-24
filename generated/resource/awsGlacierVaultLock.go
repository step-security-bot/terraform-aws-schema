package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlacierVaultLock = `{
  "block": {
    "attributes": {
      "complete_lock": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ignore_deletion_error": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "policy": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vault_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGlacierVaultLockSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlacierVaultLock), &result)
	return &result
}
