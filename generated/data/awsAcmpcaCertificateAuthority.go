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
      "key_storage_security_standard": {
        "computed": true,
        "description_kind": "plain",
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
      "revocation_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "crl_configuration": [
                "list",
                [
                  "object",
                  {
                    "custom_cname": "string",
                    "enabled": "bool",
                    "expiration_in_days": "number",
                    "s3_bucket_name": "string",
                    "s3_object_acl": "string"
                  }
                ]
              ],
              "ocsp_configuration": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool",
                    "ocsp_custom_cname": "string"
                  }
                ]
              ]
            }
          ]
        ]
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
      },
      "usage_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
