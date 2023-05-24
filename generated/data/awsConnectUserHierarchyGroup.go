package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsConnectUserHierarchyGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hierarchy_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hierarchy_path": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "level_five": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "level_four": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "level_one": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "level_three": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "level_two": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "id": "string",
                    "name": "string"
                  }
                ]
              ]
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
      "instance_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "level_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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

func AwsConnectUserHierarchyGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsConnectUserHierarchyGroup), &result)
	return &result
}
