package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEfsAccessPoint = `{
  "block": {
    "attributes": {
      "access_point_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "file_system_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "file_system_id": {
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
      "owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "posix_user": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "gid": "number",
              "secondary_gids": [
                "set",
                "number"
              ],
              "uid": "number"
            }
          ]
        ]
      },
      "root_directory": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "creation_info": [
                "list",
                [
                  "object",
                  {
                    "owner_gid": "number",
                    "owner_uid": "number",
                    "permissions": "string"
                  }
                ]
              ],
              "path": "string"
            }
          ]
        ]
      },
      "tags": {
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

func AwsEfsAccessPointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEfsAccessPoint), &result)
	return &result
}
