package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsStoragegatewaySmbFileShare = `{
  "block": {
    "attributes": {
      "access_based_enumeration": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "admin_user_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "audit_destination_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "authentication": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bucket_region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "case_sensitivity": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_storage_class": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "file_share_name": {
        "computed": true,
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
      "notification_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "object_acl": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "oplocks_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "smb_acl_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "valid_user_list": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "vpc_endpoint_dns_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "cache_attributes": {
        "block": {
          "attributes": {
            "cache_stale_timeout_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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
