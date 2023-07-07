package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsTimestreamwriteTable = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "database_name": {
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
      "table_name": {
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
      }
    },
    "block_types": {
      "magnetic_store_write_properties": {
        "block": {
          "attributes": {
            "enable_magnetic_store_writes": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "magnetic_store_rejected_data_location": {
              "block": {
                "block_types": {
                  "s3_configuration": {
                    "block": {
                      "attributes": {
                        "bucket_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "encryption_option": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "kms_key_id": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "object_key_prefix": {
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
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "retention_properties": {
        "block": {
          "attributes": {
            "magnetic_store_retention_period_in_days": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "memory_store_retention_period_in_hours": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "schema": {
        "block": {
          "block_types": {
            "composite_partition_key": {
              "block": {
                "attributes": {
                  "enforcement_in_record": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
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
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsTimestreamwriteTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsTimestreamwriteTable), &result)
	return &result
}
