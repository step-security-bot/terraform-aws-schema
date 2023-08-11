package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlueCrawler = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "classifiers": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "configuration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_configuration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "table_prefix": {
        "description_kind": "plain",
        "optional": true,
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
      "catalog_target": {
        "block": {
          "attributes": {
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
            "dlq_event_queue_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "event_queue_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tables": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "delta_target": {
        "block": {
          "attributes": {
            "connection_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_native_delta_table": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "delta_tables": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "write_manifest": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "dynamodb_target": {
        "block": {
          "attributes": {
            "path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "scan_all": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "scan_rate": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "hudi_target": {
        "block": {
          "attributes": {
            "connection_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "exclusions": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "maximum_traversal_depth": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "paths": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "iceberg_target": {
        "block": {
          "attributes": {
            "connection_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "exclusions": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "maximum_traversal_depth": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "paths": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "jdbc_target": {
        "block": {
          "attributes": {
            "connection_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "enable_additional_metadata": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "exclusions": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "lake_formation_configuration": {
        "block": {
          "attributes": {
            "account_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_lake_formation_credentials": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "lineage_configuration": {
        "block": {
          "attributes": {
            "crawler_lineage_settings": {
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
      "mongodb_target": {
        "block": {
          "attributes": {
            "connection_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "scan_all": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "recrawl_policy": {
        "block": {
          "attributes": {
            "recrawl_behavior": {
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
      "s3_target": {
        "block": {
          "attributes": {
            "connection_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dlq_event_queue_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "event_queue_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "exclusions": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sample_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "schema_change_policy": {
        "block": {
          "attributes": {
            "delete_behavior": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_behavior": {
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

func AwsGlueCrawlerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlueCrawler), &result)
	return &result
}
