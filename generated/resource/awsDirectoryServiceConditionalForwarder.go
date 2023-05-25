package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDirectoryServiceConditionalForwarder = `{
  "block": {
    "attributes": {
      "directory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dns_ips": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remote_domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDirectoryServiceConditionalForwarderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDirectoryServiceConditionalForwarder), &result)
	return &result
}
