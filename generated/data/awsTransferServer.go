package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsTransferServer = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint": {
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
      "identity_provider_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "invocation_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "logging_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "server_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsTransferServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsTransferServer), &result)
	return &result
}
