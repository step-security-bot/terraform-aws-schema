package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftClusterCredentials = `{
  "block": {
    "attributes": {
      "auto_create": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cluster_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_groups": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
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
        "description_kind": "plain",
        "required": true,
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRedshiftClusterCredentialsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftClusterCredentials), &result)
	return &result
}
