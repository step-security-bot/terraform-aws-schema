package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAmplifyApp = `{
  "block": {
    "attributes": {
      "access_token": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_branch_creation_patterns": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "basic_auth_credentials": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "build_spec": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_auto_branch_creation": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_basic_auth": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_branch_auto_build": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_branch_auto_deletion": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "environment_variables": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "iam_service_role_arn": {
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
      "oauth_token": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "platform": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "production_branch": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "branch_name": "string",
              "last_deploy_time": "string",
              "status": "string",
              "thumbnail_url": "string"
            }
          ]
        ]
      },
      "repository": {
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
      "auto_branch_creation_config": {
        "block": {
          "attributes": {
            "basic_auth_credentials": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "build_spec": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enable_auto_build": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_basic_auth": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_performance_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_pull_request_preview": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "environment_variables": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "framework": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pull_request_environment_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "stage": {
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
      "custom_rule": {
        "block": {
          "attributes": {
            "condition": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "status": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target": {
              "description_kind": "plain",
              "required": true,
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

func AwsAmplifyAppSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAmplifyApp), &result)
	return &result
}
