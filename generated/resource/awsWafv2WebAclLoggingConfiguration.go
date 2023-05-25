package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsWafv2WebAclLoggingConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_destination_configs": {
        "description": "AWS Kinesis Firehose Delivery Stream ARNs",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "resource_arn": {
        "description": "AWS WebACL ARN",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "redacted_fields": {
        "block": {
          "block_types": {
            "all_query_arguments": {
              "block": {
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "body": {
              "block": {
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "method": {
              "block": {
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "query_string": {
              "block": {
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "single_header": {
              "block": {
                "attributes": {
                  "name": {
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
            "single_query_argument": {
              "block": {
                "attributes": {
                  "name": {
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
            "uri_path": {
              "block": {
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Parts of the request to exclude from logs",
          "description_kind": "plain"
        },
        "max_items": 100,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsWafv2WebAclLoggingConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsWafv2WebAclLoggingConfiguration), &result)
	return &result
}
