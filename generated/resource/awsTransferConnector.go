package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsTransferConnector = `{
  "block": {
    "attributes": {
      "access_role": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connector_id": {
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
      "logging_role": {
        "description_kind": "plain",
        "optional": true,
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
      "url": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "as2_config": {
        "block": {
          "attributes": {
            "compression": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "encryption_algorithm": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "local_profile_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "mdn_response": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "mdn_signing_algorithm": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "message_subject": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "partner_profile_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "signing_algorithm": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "sftp_config": {
        "block": {
          "attributes": {
            "trusted_host_keys": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "user_secret_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsTransferConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsTransferConnector), &result)
	return &result
}
