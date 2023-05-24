package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDatasyncLocationHdfs = `{
  "block": {
    "attributes": {
      "agent_arns": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authentication_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "block_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kerberos_keytab": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kerberos_krb5_conf": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kerberos_principal": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_provider_uri": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replication_factor": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "simple_user": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subdirectory": {
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
      "uri": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "name_node": {
        "block": {
          "attributes": {
            "hostname": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      },
      "qop_configuration": {
        "block": {
          "attributes": {
            "data_transfer_protection": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rpc_protection": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsDatasyncLocationHdfsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDatasyncLocationHdfs), &result)
	return &result
}
