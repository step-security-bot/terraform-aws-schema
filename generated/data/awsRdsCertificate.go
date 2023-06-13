package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRdsCertificate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_override": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "customer_override_valid_till": {
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
      "latest_valid_till": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "thumbprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "valid_from": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "valid_till": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRdsCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRdsCertificate), &result)
	return &result
}
