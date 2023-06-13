package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftserverlessCredentials = `{
  "block": {
    "attributes": {
      "db_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_password": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "db_user": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "duration_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "expiration": {
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
      "workgroup_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRedshiftserverlessCredentialsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftserverlessCredentials), &result)
	return &result
}
