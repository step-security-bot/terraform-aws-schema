package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRdsExportTask = `{
  "block": {
    "attributes": {
      "export_only": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "export_task_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "failure_cause": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "iam_role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "percent_progress": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "s3_bucket_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshot_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "task_end_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "task_start_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "warning_message": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.",
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

func AwsRdsExportTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRdsExportTask), &result)
	return &result
}
