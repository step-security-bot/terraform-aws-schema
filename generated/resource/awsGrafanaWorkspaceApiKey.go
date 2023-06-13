package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGrafanaWorkspaceApiKey = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key_role": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "seconds_to_live": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "workspace_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGrafanaWorkspaceApiKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGrafanaWorkspaceApiKey), &result)
	return &result
}
