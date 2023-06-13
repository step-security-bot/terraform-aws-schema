package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDxGatewayAssociationProposal = `{
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "associated_gateway_owner_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "associated_gateway_type": {
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
        "description_kind": "plain",
        "required": true,
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

func AwsDxGatewayAssociationProposalSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDxGatewayAssociationProposal), &result)
	return &result
}
