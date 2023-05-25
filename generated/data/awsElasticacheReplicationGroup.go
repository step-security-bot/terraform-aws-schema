package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsElasticacheReplicationGroup = `{
  "block": {
    "attributes": {
      "auth_token_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "automatic_failover_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "configuration_endpoint_address": {
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
      "member_clusters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "node_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "number_cache_clusters": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "primary_endpoint_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "replication_group_description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "replication_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsElasticacheReplicationGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsElasticacheReplicationGroup), &result)
	return &result
}
