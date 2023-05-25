package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDxGatewayAssociation = `{
  "block": {
    "attributes": {
      "allowed_prefixes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "associated_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "associated_gateway_owner_account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "associated_gateway_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dx_gateway_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dx_gateway_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dx_gateway_owner_account_id": {
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
      "proposal_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpn_gateway_id": {
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
  "version": 1
}`

func AwsDxGatewayAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDxGatewayAssociation), &result)
	return &result
}
