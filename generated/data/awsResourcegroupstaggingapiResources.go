package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsResourcegroupstaggingapiResources = `{
  "block": {
    "attributes": {
      "exclude_compliant_resources": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "include_compliance_details": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_arn_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "resource_tag_mapping_list": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "compliance_details": [
                "list",
                [
                  "object",
                  {
                    "compliance_status": "bool",
                    "keys_with_noncompliant_values": [
                      "set",
                      "string"
                    ],
                    "non_compliant_keys": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "resource_arn": "string",
              "tags": [
                "map",
                "string"
              ]
            }
          ]
        ]
      },
      "resource_type_filters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "tag_filter": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 50,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsResourcegroupstaggingapiResourcesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsResourcegroupstaggingapiResources), &result)
	return &result
}
