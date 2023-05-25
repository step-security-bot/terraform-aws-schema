package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDbInstance = `{
  "block": {
    "attributes": {
      "address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "allocated_storage": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "auto_minor_version_upgrade": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_retention_period": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "ca_cert_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_class": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_instance_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "db_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_parameter_groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "db_security_groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "db_subnet_group": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled_cloudwatch_logs_exports": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "endpoint": {
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
      "hosted_zone_id": {
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
      "iops": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "license_model": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "master_username": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "monitoring_interval": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "monitoring_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "multi_az": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "option_group_memberships": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "preferred_backup_window": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "replicate_source_db": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_encrypted": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "storage_type": {
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
      "timezone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_security_groups": {
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
  "version": 1
}`

func AwsDbInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDbInstance), &result)
	return &result
}
