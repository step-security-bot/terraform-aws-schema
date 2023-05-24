package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAmplifyBranch = `{
  "block": {
    "attributes": {
      "app_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "associated_resources": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "backend_environment_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "basic_auth_credentials": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "branch_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "custom_domains": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_branch": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
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
      "enable_notification": {
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pull_request_environment_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_branch": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stage": {
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
      },
      "ttl": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAmplifyBranchSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAmplifyBranch), &result)
	return &result
}
