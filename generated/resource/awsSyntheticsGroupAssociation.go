package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSyntheticsGroupAssociation = `{
  "block": {
    "attributes": {
      "canary_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSyntheticsGroupAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSyntheticsGroupAssociation), &result)
	return &result
}
