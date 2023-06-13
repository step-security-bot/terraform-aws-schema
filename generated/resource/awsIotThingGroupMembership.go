package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIotThingGroupMembership = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "override_dynamic_group": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "thing_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "thing_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsIotThingGroupMembershipSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIotThingGroupMembership), &result)
	return &result
}
