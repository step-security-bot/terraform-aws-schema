package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsS3BucketObject = `{
  "block": {
    "attributes": {
      "body": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bucket": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "bucket_key_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "cache_control": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "content_disposition": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "content_encoding": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "content_language": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "content_length": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "content_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expiration": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expires": {
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
      "key": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "metadata": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "object_lock_legal_hold_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "object_lock_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "object_lock_retain_until_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "range": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_side_encryption": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sse_kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_class": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "version_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "website_redirect_location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsS3BucketObjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsS3BucketObject), &result)
	return &result
}
