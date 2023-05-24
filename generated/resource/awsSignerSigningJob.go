package resource

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
      "ignore_signing_job_failure": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "job_id": {
        "computed": true,
        "description_kind": "plain",
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
        "description_kind": "plain",
        "required": true,
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
    "block_types": {
      "destination": {
        "block": {
          "block_types": {
            "s3": {
              "block": {
                "attributes": {
                  "bucket": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "prefix": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "source": {
        "block": {
          "block_types": {
            "s3": {
              "block": {
                "attributes": {
                  "bucket": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "version": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
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
