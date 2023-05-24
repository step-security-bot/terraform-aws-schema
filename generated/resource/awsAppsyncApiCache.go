package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppsyncApiCache = `{
  "block": {
    "attributes": {
      "api_caching_behavior": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "api_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "at_rest_encryption_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_encryption_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ttl": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAppsyncApiCacheSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppsyncApiCache), &result)
	return &result
}
