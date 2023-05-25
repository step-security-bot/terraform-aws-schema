package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsOpsworksApplication = `{
  "block": {
    "attributes": {
      "auto_bundle_on_deploy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "aws_flow_ruby_settings": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_source_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_source_database_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_source_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "document_root": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domains": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "enable_ssl": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rails_env": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "short_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stack_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "app_source": {
        "block": {
          "attributes": {
            "password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "revision": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssh_key": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "username": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "environment": {
        "block": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "secure": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "ssl_configuration": {
        "block": {
          "attributes": {
            "certificate": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "chain": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "private_key": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsOpsworksApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsOpsworksApplication), &result)
	return &result
}
