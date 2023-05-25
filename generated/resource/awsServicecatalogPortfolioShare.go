package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsServicecatalogPortfolioShare = `{
  "block": {
    "attributes": {
      "accept_language": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "accepted": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "portfolio_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "share_tag_options": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "wait_for_acceptance": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsServicecatalogPortfolioShareSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsServicecatalogPortfolioShare), &result)
	return &result
}
