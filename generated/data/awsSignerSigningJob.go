package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSignerSigningJob = `{
  "block": {
    "attributes": {
      "completed_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
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
      "job_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "job_invoker": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "job_owner": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "platform_display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "platform_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "profile_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "profile_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "requested_by": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "revocation_record": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "reason": "string",
              "revoked_at": "string",
              "revoked_by": "string"
            }
          ]
        ]
      },
      "signature_expires_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "signed_object": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "s3": [
                "list",
                [
                  "object",
                  {
                    "bucket": "string",
                    "key": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "source": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "s3": [
                "list",
                [
                  "object",
                  {
                    "bucket": "string",
                    "key": "string",
                    "version": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsSignerSigningJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSignerSigningJob), &result)
	return &result
}
