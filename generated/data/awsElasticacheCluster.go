package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsElasticacheCluster = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cache_nodes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address": "string",
              "availability_zone": "string",
              "id": "string",
              "outpost_arn": "string",
              "port": "number"
            }
          ]
        ]
      },
      "cluster_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "configuration_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
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
      "ip_discovery": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "log_delivery_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "destination": "string",
              "destination_type": "string",
              "log_format": "string",
              "log_type": "string"
            }
          ]
        ]
      },
      "maintenance_window": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "node_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "notification_topic_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "num_cache_nodes": {
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
      "preferred_outpost_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "replication_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsElasticacheClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsElasticacheCluster), &result)
	return &result
}
