package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDirectoryServiceTrust = `{
  "block": {
    "attributes": {
      "conditional_forwarder_ip_addrs": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "created_date_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "delete_associated_conditional_forwarder": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "directory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_date_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "remote_domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "selective_auth": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state_last_updated_date_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "trust_direction": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "trust_password": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "trust_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "trust_state_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "trust_type": {
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

func AwsDirectoryServiceTrustSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDirectoryServiceTrust), &result)
	return &result
}
