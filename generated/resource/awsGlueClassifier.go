package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlueClassifier = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "csv_classifier": {
        "block": {
          "attributes": {
            "allow_single_column": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "contains_header": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "custom_datatype_configured": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "custom_datatypes": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "delimiter": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "disable_value_trimming": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "header": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "quote_symbol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "grok_classifier": {
        "block": {
          "attributes": {
            "classification": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "custom_patterns": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "grok_pattern": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "json_classifier": {
        "block": {
          "attributes": {
            "json_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "xml_classifier": {
        "block": {
          "attributes": {
            "classification": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "row_tag": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGlueClassifierSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlueClassifier), &result)
	return &result
}
