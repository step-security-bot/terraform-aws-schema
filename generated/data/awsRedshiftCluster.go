package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRedshiftCluster = `{
  "block": {
    "attributes": {
      "allow_version_upgrade": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "aqua_configuration_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "automated_snapshot_retention_period": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_relocation_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "bucket_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_namespace_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_nodes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "node_role": "string",
              "private_ip_address": "string",
              "public_ip_address": "string"
            }
          ]
        ]
      },
      "cluster_parameter_group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_public_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_revision_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_subnet_group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "database_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_iam_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "elastic_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_logging": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "encrypted": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enhanced_vpc_routing": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "iam_roles": {
        "computed": true,
        "description_kind": "plain",
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
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "log_destination_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "log_exports": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "maintenance_track_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "manual_snapshot_retention_period": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "master_username": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "node_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "number_of_nodes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "publicly_accessible": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "s3_key_prefix": {
        "computed": true,
        "description_kind": "plain",
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
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRedshiftClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRedshiftCluster), &result)
	return &result
}
