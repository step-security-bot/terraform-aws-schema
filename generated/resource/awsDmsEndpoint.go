package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDmsEndpoint = `{
  "block": {
    "attributes": {
      "certificate_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "engine_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "extra_connection_attributes": {
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
      "kms_key_arn": {
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
      "port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "server_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_access_role": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ssl_mode": {
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
      "username": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "elasticsearch_settings": {
        "block": {
          "attributes": {
            "endpoint_uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "error_retry_duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "full_load_error_percentage": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "service_access_role_arn": {
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
      "kafka_settings": {
        "block": {
          "attributes": {
            "broker": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "topic": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "kinesis_settings": {
        "block": {
          "attributes": {
            "message_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_access_role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "stream_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "mongodb_settings": {
        "block": {
          "attributes": {
            "auth_mechanism": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "auth_source": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "auth_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "docs_to_investigate": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "extract_doc_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "nesting_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "s3_settings": {
        "block": {
          "attributes": {
            "bucket_folder": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bucket_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "compression_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "csv_delimiter": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "csv_row_delimiter": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "external_table_definition": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_access_role_arn": {
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

func AwsDmsEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDmsEndpoint), &result)
	return &result
}
