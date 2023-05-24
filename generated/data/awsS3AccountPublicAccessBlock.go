package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3AccountPublicAccessBlock = `{
  "block": {
    "attributes": {
      "account_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "block_public_acls": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "block_public_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ignore_public_acls": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "restrict_public_buckets": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3AccountPublicAccessBlockSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3AccountPublicAccessBlock), &result)
	return &result
}
