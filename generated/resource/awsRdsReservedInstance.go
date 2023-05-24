package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRdsReservedInstance = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "currency_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_class": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "duration": {
        "computed": true,
        "description_kind": "plain",
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
      "instance_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "lease_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "multi_az": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "offering_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "offering_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "product_description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "recurring_charges": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "recurring_charge_amount": "number",
              "recurring_charge_frequency": "string"
            }
          ]
        ]
      },
      "reservation_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "usage_price": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRdsReservedInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRdsReservedInstance), &result)
	return &result
}
