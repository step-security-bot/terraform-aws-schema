package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAuditmanagerControl = `{
  "block": {
    "attributes": {
      "action_plan_instructions": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "action_plan_title": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "testing_information": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "control_mapping_sources": {
        "block": {
          "attributes": {
            "source_description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source_frequency": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source_set_up_option": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "troubleshooting_text": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "source_keyword": {
              "block": {
                "attributes": {
                  "keyword_input_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "keyword_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAuditmanagerControlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAuditmanagerControl), &result)
	return &result
}
