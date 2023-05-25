package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAcmpcaCertificateAuthorityCertificate = `{
  "block": {
    "attributes": {
      "certificate": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate_authority_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate_chain": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAcmpcaCertificateAuthorityCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAcmpcaCertificateAuthorityCertificate), &result)
	return &result
}
