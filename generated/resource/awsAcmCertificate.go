package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAcmCertificate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_authority_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_body": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_chain": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_validation_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "domain_name": "string",
              "resource_record_name": "string",
              "resource_record_type": "string",
              "resource_record_value": "string"
            }
          ]
        ]
      },
      "early_renewal_duration": {
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
      "key_algorithm": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "not_after": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "not_before": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pending_renewal": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "private_key": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "renewal_eligibility": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "renewal_summary": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "renewal_status": "string",
              "renewal_status_reason": "string",
              "updated_at": "string"
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subject_alternative_names": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
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
      },
      "validation_emails": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "validation_method": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "options": {
        "block": {
          "attributes": {
            "certificate_transparency_logging_preference": {
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
      "validation_option": {
        "block": {
          "attributes": {
            "domain_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "validation_domain": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAcmCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAcmCertificate), &result)
	return &result
}
