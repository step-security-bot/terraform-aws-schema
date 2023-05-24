package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlueDevEndpoint = `{
  "block": {
    "attributes": {
      "arguments": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "extra_jars_s3_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "extra_python_libs_s3_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failure_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "glue_version": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "number_of_nodes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "number_of_workers": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "private_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_keys": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_configuration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_id": {
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
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "worker_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "yarn_endpoint_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "zeppelin_remote_spark_interpreter_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGlueDevEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlueDevEndpoint), &result)
	return &result
}
