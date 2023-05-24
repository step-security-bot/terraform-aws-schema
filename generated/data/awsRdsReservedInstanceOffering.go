package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRdsReservedInstanceOffering = `{
  "block": {
    "attributes": {
      "currency_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_class": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "duration": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "fixed_price": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_az": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "offering_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "offering_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_description": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRdsReservedInstanceOfferingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRdsReservedInstanceOffering), &result)
	return &result
}
