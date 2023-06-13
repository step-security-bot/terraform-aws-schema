package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsKmsKey = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cloud_hsm_cluster_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_key_store_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_master_key_spec": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "expiration_model": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "grant_tokens": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key_manager": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_spec": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_usage": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "multi_region": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "multi_region_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "multi_region_key_type": "string",
              "primary_key": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "region": "string"
                  }
                ]
              ],
              "replica_keys": [
                "list",
                [
                  "object",
                  {
                    "arn": "string",
                    "region": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "origin": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pending_deletion_window_in_days": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "valid_to": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "xks_key_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsKmsKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsKmsKey), &result)
	return &result
}
