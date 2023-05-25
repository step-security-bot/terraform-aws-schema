package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGuarddutyPublishingDestination = `{
  "block": {
    "attributes": {
      "destination_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "detector_id": {
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
      "kms_key_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGuarddutyPublishingDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGuarddutyPublishingDestination), &result)
	return &result
}
