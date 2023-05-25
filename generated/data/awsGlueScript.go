package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlueScript = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "language": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "python_script": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "scala_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "dag_edge": {
        "block": {
          "attributes": {
            "source": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_parameter": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "dag_node": {
        "block": {
          "attributes": {
            "id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "line_number": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "node_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "args": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "param": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGlueScriptSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlueScript), &result)
	return &result
}
