package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRdsOrderableDbInstance = `{
  "block": {
    "attributes": {
      "availability_zone_group": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "engine": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "engine_version": {
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
      "instance_class": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "license_model": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_iops_per_db_instance": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_iops_per_gib": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "max_storage_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "min_iops_per_db_instance": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "min_iops_per_gib": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "min_storage_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "multi_az_capable": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "outpost_capable": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "preferred_engine_versions": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "preferred_instance_classes": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "read_replica_capable": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "storage_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "supported_engine_modes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "supported_network_types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "supports_enhanced_monitoring": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "supports_global_databases": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "supports_iam_database_authentication": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "supports_iops": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "supports_kerberos_authentication": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "supports_performance_insights": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "supports_storage_autoscaling": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "supports_storage_encryption": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "vpc": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRdsOrderableDbInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRdsOrderableDbInstance), &result)
	return &result
}
