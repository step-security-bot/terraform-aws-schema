package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsConnectUserHierarchyStructure = `{
  "block": {
    "attributes": {
      "hierarchy_structure": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsConnectUserHierarchyStructureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsConnectUserHierarchyStructure), &result)
	return &result
}
