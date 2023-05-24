package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAcmpcaCertificate = `{
  "block": {
    "attributes": {
      "api_passthrough": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_authority_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate_chain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_signing_request": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "signing_algorithm": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "validity": {
        "block": {
          "attributes": {
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
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
  "version": 0
}`

func AwsAcmpcaCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAcmpcaCertificate), &result)
	return &result
}
