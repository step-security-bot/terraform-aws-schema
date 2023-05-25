package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSnsTopic = `{
  "block": {
    "attributes": {
      "application_failure_feedback_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "application_success_feedback_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "application_success_feedback_sample_rate": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "content_based_deduplication": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "delivery_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fifo_topic": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "firehose_failure_feedback_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "firehose_success_feedback_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "firehose_success_feedback_sample_rate": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "http_failure_feedback_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "http_success_feedback_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "http_success_feedback_sample_rate": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_master_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lambda_failure_feedback_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lambda_success_feedback_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lambda_success_feedback_sample_rate": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sqs_failure_feedback_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sqs_success_feedback_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sqs_success_feedback_sample_rate": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSnsTopicSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSnsTopic), &result)
	return &result
}
