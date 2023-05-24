package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftAuthenticationProfile = `{
  "block": {
    "attributes": {
      "authentication_profile_content": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "authentication_profile_name": {
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

func AwsRedshiftAuthenticationProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftAuthenticationProfile), &result)
	return &result
}
