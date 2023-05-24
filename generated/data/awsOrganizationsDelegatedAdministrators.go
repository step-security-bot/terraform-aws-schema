package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOrganizationsDelegatedAdministrators = `{
  "block": {
    "attributes": {
      "delegated_administrators": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "arn": "string",
              "delegation_enabled_date": "string",
              "email": "string",
              "id": "string",
              "joined_method": "string",
              "joined_timestamp": "string",
              "name": "string",
              "status": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_principal": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOrganizationsDelegatedAdministratorsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOrganizationsDelegatedAdministrators), &result)
	return &result
}
