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
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
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
