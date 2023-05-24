package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsEcrReplicationConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "registry_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "replication_configuration": {
        "block": {
          "block_types": {
            "rule": {
              "block": {
                "block_types": {
                  "destination": {
                    "block": {
                      "attributes": {
                        "region": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "registry_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 25,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "repository_filter": {
                    "block": {
                      "attributes": {
                        "filter": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "filter_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 100,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 10,
              "min_items": 1,
              "nesting_mode": "list"
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

func AwsEcrReplicationConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsEcrReplicationConfiguration), &result)
	return &result
}
