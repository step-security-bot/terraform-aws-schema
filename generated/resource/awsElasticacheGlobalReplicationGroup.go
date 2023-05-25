package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsElasticacheGlobalReplicationGroup = `{
  "block": {
    "attributes": {
      "actual_engine_version": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "at_rest_encryption_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "auth_token_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "cache_node_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "engine": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version_actual": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "global_replication_group_description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "global_replication_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "global_replication_group_id_suffix": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "primary_replication_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_encryption_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsElasticacheGlobalReplicationGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsElasticacheGlobalReplicationGroup), &result)
	return &result
}
