package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCodepipelineCustomActionType = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "category": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provider_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "configuration_property": {
        "block": {
          "attributes": {
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "queryable": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "required": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "secret": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 10,
        "nesting_mode": "list"
      },
      "input_artifact_details": {
        "block": {
          "attributes": {
            "maximum_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "minimum_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "output_artifact_details": {
        "block": {
          "attributes": {
            "maximum_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "minimum_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "settings": {
        "block": {
          "attributes": {
            "entity_url_template": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "execution_url_template": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "revision_url_template": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "third_party_configuration_url": {
              "description_kind": "plain",
              "optional": true,
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

func AwsCodepipelineCustomActionTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCodepipelineCustomActionType), &result)
	return &result
}
