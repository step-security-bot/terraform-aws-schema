package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLambdaFunction = `{
  "block": {
    "attributes": {
      "architectures": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "code_signing_config_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dead_letter_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "target_arn": "string"
            }
          ]
        ]
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "environment": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "variables": [
                "map",
                "string"
              ]
            }
          ]
        ]
      },
      "ephemeral_storage": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "size": "number"
            }
          ]
        ]
      },
      "file_system_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "local_mount_path": "string"
            }
          ]
        ]
      },
      "function_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "handler": {
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
      "image_uri": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "invoke_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "layers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "memory_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "qualified_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "qualified_invoke_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "qualifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reserved_concurrent_executions": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "runtime": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "signing_job_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "signing_profile_version_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_code_hash": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_code_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "timeout": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tracing_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "mode": "string"
            }
          ]
        ]
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "security_group_ids": [
                "set",
                "string"
              ],
              "subnet_ids": [
                "set",
                "string"
              ],
              "vpc_id": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsLambdaFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLambdaFunction), &result)
	return &result
}
