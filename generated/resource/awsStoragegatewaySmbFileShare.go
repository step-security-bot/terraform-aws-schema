package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsStoragegatewaySmbFileShare = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authentication": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_storage_class": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fileshare_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "guess_mime_type_enabled": {
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
      "invalid_user_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "kms_encrypted": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "kms_key_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "object_acl": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "path": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "read_only": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "requester_pays": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "role_arn": {
        "description_kind": "plain",
        "required": true,
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
      "valid_user_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
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

func AwsStoragegatewaySmbFileShareSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsStoragegatewaySmbFileShare), &result)
	return &result
}
