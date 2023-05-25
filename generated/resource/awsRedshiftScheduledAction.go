package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftScheduledAction = `{
  "block": {
    "attributes": {
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "end_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "iam_role": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "start_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "target_action": {
        "block": {
          "block_types": {
            "pause_cluster": {
              "block": {
                "attributes": {
                  "cluster_identifier": {
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
            "resize_cluster": {
              "block": {
                "attributes": {
                  "classic": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "cluster_identifier": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "cluster_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "node_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "number_of_nodes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "resume_cluster": {
              "block": {
                "attributes": {
                  "cluster_identifier": {
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
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRedshiftScheduledActionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftScheduledAction), &result)
	return &result
}
