package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudtrail = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cloud_watch_logs_group_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cloud_watch_logs_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_log_file_validation": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_logging": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "home_region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "include_global_service_events": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_multi_region_trail": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_organization_trail": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "kms_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_bucket_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_key_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sns_topic_name": {
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
      }
    },
    "block_types": {
      "event_selector": {
        "block": {
          "attributes": {
            "include_management_events": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "read_write_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "data_resource": {
              "block": {
                "attributes": {
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "values": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 5,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudtrailSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudtrail), &result)
	return &result
}
