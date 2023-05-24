package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLightsailDatabase = `{
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
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backup_retention_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "blueprint_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "bundle_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ca_certificate_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cpu_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disk_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "final_snapshot_name": {
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
      "master_database_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "master_endpoint_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "master_endpoint_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "master_password": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "master_username": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "preferred_backup_window": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "publicly_accessible": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ram_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "relational_database_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secondary_availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "skip_final_snapshot": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "support_code": {
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
      "tags_all": {
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

func AwsLightsailDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLightsailDatabase), &result)
	return &result
}
