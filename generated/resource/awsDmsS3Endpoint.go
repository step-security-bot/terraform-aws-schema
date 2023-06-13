package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDmsS3Endpoint = `{
  "block": {
    "attributes": {
      "add_column_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "add_trailing_padding_character": {
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
        "required": true,
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
      "certificate_arn": {
        "computed": true,
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
      "date_partition_timezone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "detach_target_on_lob_lookup_failure_parquet": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "engine_display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expected_bucket_owner": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "external_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "external_table_definition": {
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
      "ignore_header_rows": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "include_op_for_full_load": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "kms_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
        "required": true,
        "type": "string"
      },
      "ssl_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
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
    "block_types": {
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

func AwsDmsS3EndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDmsS3Endpoint), &result)
	return &result
}
