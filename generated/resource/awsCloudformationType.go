package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudformationType = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_version_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deprecated_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "documentation_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "execution_role_arn": {
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
      "is_default_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "provisioning_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "schema": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "schema_handler_package": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "version_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "visibility": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "logging_config": {
        "block": {
          "attributes": {
            "log_group_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "log_role_arn": {
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
  "version": 0
}`

func AwsCloudformationTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudformationType), &result)
	return &result
}
