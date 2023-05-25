package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsServicecatalogProvisioningArtifact = `{
  "block": {
    "attributes": {
      "accept_language": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "active": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "created_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disable_template_validation": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "guidance": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "product_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_physical_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_url": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsServicecatalogProvisioningArtifactSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsServicecatalogProvisioningArtifact), &result)
	return &result
}
