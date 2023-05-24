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
      "secrets_manager_access_role_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secrets_manager_arn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "tags_all": {
        "computed": true,
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
            "include_control_details": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_null_and_empty": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_partition_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_table_alter_operations": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_transaction_details": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "message_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "message_max_bytes": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "no_hex_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "partition_include_schema_table": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "sasl_password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "sasl_username": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_protocol": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_ca_certificate_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_client_certificate_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_client_key_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_client_key_password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
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
            "include_control_details": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_null_and_empty": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_partition_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_table_alter_operations": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_transaction_details": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "message_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "partition_include_schema_table": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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
      "redis_settings": {
        "block": {
          "attributes": {
            "auth_password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "auth_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "auth_user_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "server_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ssl_ca_certificate_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_security_protocol": {
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
      "redshift_settings": {
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
            "encryption_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "server_side_encryption_kms_key_id": {
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
      },
      "s3_settings": {
        "block": {
          "attributes": {
            "add_column_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
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
            "canned_acl_for_objects": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cdc_inserts_and_updates": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "cdc_inserts_only": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "cdc_max_batch_interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "cdc_min_file_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "cdc_path": {
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
            "csv_no_sup_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "csv_null_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "csv_row_delimiter": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_format": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_page_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "date_partition_delimiter": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "date_partition_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "date_partition_sequence": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dict_page_size_limit": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enable_statistics": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "encoding_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encryption_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "external_table_definition": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ignore_header_rows": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ignore_headers_row": {
              "description": "This setting has no effect, is deprecated, and will be removed in a future version",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "include_op_for_full_load": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "max_file_size": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "parquet_timestamp_in_millisecond": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "parquet_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "preserve_transactions": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "rfc_4180": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "row_group_length": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "server_side_encryption_kms_key_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_access_role_arn": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "timestamp_column_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_csv_no_sup_value": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "use_task_start_time_for_full_load_timestamp": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "This argument is deprecated and will be removed in a future version; use aws_dms_s3_endpoint instead",
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
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
