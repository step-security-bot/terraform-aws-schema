package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAccountPrimaryContact = `{
  "block": {
    "attributes": {
      "account_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "address_line_1": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "address_line_2": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "address_line_3": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "city": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "company_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "country_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "district_or_county": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "full_name": {
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
      "phone_number": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "postal_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state_or_region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "website_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAccountPrimaryContactSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAccountPrimaryContact), &result)
	return &result
}
