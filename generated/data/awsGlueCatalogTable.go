package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsGlueCatalogTable = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "catalog_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "partition_index": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "index_name": "string",
              "index_status": "string",
              "keys": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "partition_keys": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "comment": "string",
              "name": "string",
              "type": "string"
            }
          ]
        ]
      },
      "query_as_of_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retention": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "storage_descriptor": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bucket_columns": [
                "list",
                "string"
              ],
              "columns": [
                "list",
                [
                  "object",
                  {
                    "comment": "string",
                    "name": "string",
                    "parameters": [
                      "map",
                      "string"
                    ],
                    "type": "string"
                  }
                ]
              ],
              "compressed": "bool",
              "input_format": "string",
              "location": "string",
              "number_of_buckets": "number",
              "output_format": "string",
              "parameters": [
                "map",
                "string"
              ],
              "schema_reference": [
                "list",
                [
                  "object",
                  {
                    "schema_id": [
                      "list",
                      [
                        "object",
                        {
                          "registry_name": "string",
                          "schema_arn": "string",
                          "schema_name": "string"
                        }
                      ]
                    ],
                    "schema_version_id": "string",
                    "schema_version_number": "number"
                  }
                ]
              ],
              "ser_de_info": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "parameters": [
                      "map",
                      "string"
                    ],
                    "serialization_library": "string"
                  }
                ]
              ],
              "skewed_info": [
                "list",
                [
                  "object",
                  {
                    "skewed_column_names": [
                      "list",
                      "string"
                    ],
                    "skewed_column_value_location_maps": [
                      "map",
                      "string"
                    ],
                    "skewed_column_values": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "sort_columns": [
                "list",
                [
                  "object",
                  {
                    "column": "string",
                    "sort_order": "number"
                  }
                ]
              ],
              "stored_as_sub_directories": "bool"
            }
          ]
        ]
      },
      "table_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_table": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "catalog_id": "string",
              "database_name": "string",
              "name": "string"
            }
          ]
        ]
      },
      "transaction_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "view_expanded_text": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "view_original_text": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsGlueCatalogTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsGlueCatalogTable), &result)
	return &result
}
