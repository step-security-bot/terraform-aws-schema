package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlueConnection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "catalog_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_properties": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": [
          "map",
          "string"
        ]
      },
      "connection_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "match_criteria": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "physical_connection_requirements": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "availability_zone": "string",
              "security_group_id_list": [
                "set",
                "string"
              ],
              "subnet_id": "string"
            }
          ]
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGlueConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlueConnection), &result)
	return &result
}
