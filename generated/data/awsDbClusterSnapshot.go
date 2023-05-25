package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDbClusterSnapshot = `{
  "block": {
    "attributes": {
      "allocated_storage": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "availability_zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "db_cluster_identifier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_cluster_snapshot_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_snapshot_identifier": {
        "description_kind": "plain",
        "optional": true,
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
      "include_public": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "include_shared": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "most_recent": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "snapshot_create_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_db_cluster_snapshot_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_encrypted": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDbClusterSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDbClusterSnapshot), &result)
	return &result
}
