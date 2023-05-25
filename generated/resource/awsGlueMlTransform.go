package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlueMlTransform = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "glue_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "label_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_capacity": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_retries": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "number_of_workers": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schema": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "data_type": "string",
              "name": "string"
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
      "timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "worker_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "input_record_tables": {
        "block": {
          "attributes": {
            "catalog_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "connection_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "database_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "table_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "parameters": {
        "block": {
          "attributes": {
            "transform_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "find_matches_parameters": {
              "block": {
                "attributes": {
                  "accuracy_cost_trade_off": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "enforce_provided_labels": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "precision_recall_trade_off": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "primary_key_column_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGlueMlTransformSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlueMlTransform), &result)
	return &result
}
