package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3AccountPublicAccessBlock = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "block_public_acls": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "block_public_policy": {
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
      "ignore_public_acls": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "restrict_public_buckets": {
        "description_kind": "plain",
        "optional": true,
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
