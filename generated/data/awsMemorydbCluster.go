package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMemorydbCluster = `{
  "block": {
    "attributes": {
      "acl_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_minor_version_upgrade": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "cluster_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address": "string",
              "port": "number"
            }
          ]
        ]
      },
      "data_tiering": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine_patch_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "final_snapshot_name": {
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
      "kms_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "maintenance_window": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "num_replicas_per_shard": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "num_shards": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "parameter_group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "shards": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "name": "string",
              "nodes": [
                "set",
                [
                  "object",
                  {
                    "availability_zone": "string",
                    "create_time": "string",
                    "endpoint": [
                      "list",
                      [
                        "object",
                        {
                          "address": "string",
                          "port": "number"
                        }
                      ]
                    ],
                    "name": "string"
                  }
                ]
              ],
              "num_nodes": "number",
              "slots": "string"
            }
          ]
        ]
      },
      "snapshot_retention_limit": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "snapshot_window": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sns_topic_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_group_name": {
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
      "tls_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsMemorydbClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMemorydbCluster), &result)
	return &result
}
