package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsElasticacheReplicationGroup = `{
  "block": {
    "attributes": {
      "apply_immediately": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "at_rest_encryption_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auth_token": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "auto_minor_version_upgrade": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "automatic_failover_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "availability_zones": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "cluster_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "configuration_endpoint_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_tiering_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_version_actual": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "final_snapshot_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "global_replication_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "maintenance_window": {
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
      "multi_az_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "node_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notification_topic_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "num_cache_clusters": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "num_node_groups": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "number_cache_clusters": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "parameter_group_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "preferred_cache_cluster_azs": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "primary_endpoint_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "reader_endpoint_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "replicas_per_node_group": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "replication_group_description": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replication_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "security_group_names": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "snapshot_arns": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "snapshot_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshot_retention_limit": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "snapshot_window": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subnet_group_name": {
        "computed": true,
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
      "transit_encryption_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "user_group_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "cluster_mode": {
        "block": {
          "attributes": {
            "num_node_groups": {
              "computed": true,
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "replicas_per_node_group": {
              "computed": true,
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "log_delivery_configuration": {
        "block": {
          "attributes": {
            "destination": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "destination_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "log_format": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "log_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 2,
        "nesting_mode": "set"
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
  "version": 1
}`

func AwsElasticacheReplicationGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsElasticacheReplicationGroup), &result)
	return &result
}