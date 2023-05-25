package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsMskCluster = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_sasl_iam": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_sasl_scram": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_brokers_tls": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kafka_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "number_of_broker_nodes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "zookeeper_connect_string": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsMskClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsMskCluster), &result)
	return &result
}
