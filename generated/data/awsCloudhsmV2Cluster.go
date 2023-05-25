package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudhsmV2Cluster = `{
  "block": {
    "attributes": {
      "cluster_certificates": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "aws_hardware_certificate": "string",
              "cluster_certificate": "string",
              "cluster_csr": "string",
              "hsm_certificate": "string",
              "manufacturer_hardware_certificate": "string"
            }
          ]
        ]
      },
      "cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_state": {
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
      "security_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
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

func AwsCloudhsmV2ClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudhsmV2Cluster), &result)
	return &result
}
