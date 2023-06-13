package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRoute53KeySigningKey = `{
  "block": {
    "attributes": {
      "digest_algorithm_mnemonic": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "digest_algorithm_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "digest_value": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dnskey_record": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ds_record": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "flag": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "hosted_zone_id": {
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
      "key_management_service_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key_tag": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "public_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "signing_algorithm_mnemonic": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "signing_algorithm_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRoute53KeySigningKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRoute53KeySigningKey), &result)
	return &result
}
