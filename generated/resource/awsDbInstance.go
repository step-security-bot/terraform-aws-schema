package resource

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
        "optional": true,
        "type": "number"
      },
      "allow_major_version_upgrade": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "apply_immediately": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_minor_version_upgrade": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backup_retention_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "backup_window": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ca_cert_identifier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "character_set_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "copy_tags_to_snapshot": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "customer_owned_ip_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "db_subnet_group_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delete_automated_backups": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "deletion_protection": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "domain": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_iam_role_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled_cloudwatch_logs_exports": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
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
      "hosted_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "iam_database_authentication_enabled": {
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
      "identifier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identifier_prefix": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_class": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "iops": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "latest_restorable_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "license_model": {
        "computed": true,
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
      "max_allocated_storage": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "monitoring_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "monitoring_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_az": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nchar_character_set_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "option_group_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameter_group_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "performance_insights_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "performance_insights_kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "performance_insights_retention_period": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "publicly_accessible": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "replica_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replicas": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "replicate_source_db": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_names": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "skip_final_snapshot": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "snapshot_identifier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_encrypted": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "storage_type": {
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
      "timezone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "username": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "restore_to_point_in_time": {
        "block": {
          "attributes": {
            "restore_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_db_instance_identifier": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_dbi_resource_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_latest_restorable_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "s3_import": {
        "block": {
          "attributes": {
            "bucket_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "bucket_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ingestion_role": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_engine": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_engine_version": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
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
  "version": 1
}`

func AwsDbInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDbInstance), &result)
	return &result
}
