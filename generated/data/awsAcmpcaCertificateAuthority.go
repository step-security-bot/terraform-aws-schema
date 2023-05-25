package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAcmpcaCertificateAuthority = `{
  "block": {
    "attributes": {
      "arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_chain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_signing_request": {
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
      "serial": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "revocation_configuration": {
        "block": {
          "block_types": {
            "crl_configuration": {
              "block": {
                "attributes": {
                  "custom_cname": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "expiration_in_days": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "s3_bucket_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_object_acl": {
                    "computed": true,
                    "description_kind": "plain",
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
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsAcmpcaCertificateAuthoritySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAcmpcaCertificateAuthority), &result)
	return &result
}
