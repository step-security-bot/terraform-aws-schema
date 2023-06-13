package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsServicecatalogProvisionedProduct = `{
  "block": {
    "attributes": {
      "accept_language": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cloudwatch_dashboard_names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "created_time": {
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
      "ignore_errors": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "last_provisioning_record_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_record_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_successful_provisioning_record_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "launch_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification_arns": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "outputs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "description": "string",
              "key": "string",
              "value": "string"
            }
          ]
        ]
      },
      "path_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "path_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "product_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "product_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "provisioning_artifact_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "provisioning_artifact_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retain_physical_resources": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
        "computed": true,
        "description_kind": "plain",
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
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "provisioning_parameters": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "use_previous_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "stack_set_provisioning_preferences": {
        "block": {
          "attributes": {
            "accounts": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "failure_tolerance_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "failure_tolerance_percentage": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_concurrency_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_concurrency_percentage": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "regions": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsServicecatalogProvisionedProductSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsServicecatalogProvisionedProduct), &result)
	return &result
}
