package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNeptuneOrderableDbInstance = `{
  "block": {
    "attributes": {
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
        "optional": true,
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
        "type": "string"
      },
      "supports_enhanced_monitoring": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_iam_database_authentication": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_iops": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_performance_insights": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_storage_encryption": {
        "computed": true,
        "description_kind": "plain",
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

func AwsNeptuneOrderableDbInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNeptuneOrderableDbInstance), &result)
	return &result
}
