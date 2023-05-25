package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSsmAssociation = `{
  "block": {
    "attributes": {
      "apply_only_at_cron_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "association_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "automation_target_parameter_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compliance_severity": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "document_version": {
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
      "instance_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_concurrency": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_errors": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "schedule_expression": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "output_location": {
        "block": {
          "attributes": {
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
            "s3_region": {
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
      "targets": {
        "block": {
          "attributes": {
            "key": {
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
        "max_items": 5,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsSsmAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSsmAssociation), &result)
	return &result
}
