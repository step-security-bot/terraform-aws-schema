package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsQuicksightDataSet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "column_groups": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "geo_spatial_column_group": [
                "list",
                [
                  "object",
                  {
                    "columns": [
                      "list",
                      "string"
                    ],
                    "country_code": "string",
                    "name": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "data_set_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_set_usage_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "disable_use_as_direct_query_source": "bool",
              "disable_use_as_imported_source": "bool"
            }
          ]
        ]
      },
      "field_folders": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "columns": [
                "list",
                "string"
              ],
              "description": "string",
              "field_folders_id": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "import_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "logical_table_map": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "alias": "string",
              "data_transforms": [
                "list",
                [
                  "object",
                  {
                    "cast_column_type_operation": [
                      "list",
                      [
                        "object",
                        {
                          "column_name": "string",
                          "format": "string",
                          "new_column_type": "string"
                        }
                      ]
                    ],
                    "create_columns_operation": [
                      "list",
                      [
                        "object",
                        {
                          "columns": [
                            "list",
                            [
                              "object",
                              {
                                "column_id": "string",
                                "column_name": "string",
                                "expression": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "filter_operation": [
                      "list",
                      [
                        "object",
                        {
                          "condition_expression": "string"
                        }
                      ]
                    ],
                    "project_operation": [
                      "list",
                      [
                        "object",
                        {
                          "projected_columns": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "rename_column_operation": [
                      "list",
                      [
                        "object",
                        {
                          "column_name": "string",
                          "new_column_name": "string"
                        }
                      ]
                    ],
                    "tag_column_operation": [
                      "list",
                      [
                        "object",
                        {
                          "column_name": "string",
                          "tags": [
                            "list",
                            [
                              "object",
                              {
                                "column_description": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "text": "string"
                                    }
                                  ]
                                ],
                                "column_geographic_role": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "untag_column_operation": [
                      "list",
                      [
                        "object",
                        {
                          "column_name": "string",
                          "tag_names": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "logical_table_map_id": "string",
              "source": [
                "list",
                [
                  "object",
                  {
                    "data_set_arn": "string",
                    "join_instruction": [
                      "list",
                      [
                        "object",
                        {
                          "left_join_key_properties": [
                            "list",
                            [
                              "object",
                              {
                                "unique_key": "bool"
                              }
                            ]
                          ],
                          "left_operand": "string",
                          "on_clause": "string",
                          "right_join_key_properties": [
                            "list",
                            [
                              "object",
                              {
                                "unique_key": "bool"
                              }
                            ]
                          ],
                          "right_operand": "string",
                          "type": "string"
                        }
                      ]
                    ],
                    "physical_table_id": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "permissions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "actions": [
                "set",
                "string"
              ],
              "principal": "string"
            }
          ]
        ]
      },
      "physical_table_map": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "custom_sql": [
                "list",
                [
                  "object",
                  {
                    "columns": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "type": "string"
                        }
                      ]
                    ],
                    "data_source_arn": "string",
                    "name": "string",
                    "sql_query": "string"
                  }
                ]
              ],
              "physical_table_map_id": "string",
              "relational_table": [
                "list",
                [
                  "object",
                  {
                    "catalog": "string",
                    "data_source_arn": "string",
                    "input_columns": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "type": "string"
                        }
                      ]
                    ],
                    "name": "string",
                    "schema": "string"
                  }
                ]
              ],
              "s3_source": [
                "list",
                [
                  "object",
                  {
                    "data_source_arn": "string",
                    "input_columns": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "type": "string"
                        }
                      ]
                    ],
                    "upload_settings": [
                      "list",
                      [
                        "object",
                        {
                          "contains_header": "bool",
                          "delimiter": "string",
                          "format": "string",
                          "start_from_row": "number",
                          "text_qualifier": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "row_level_permission_data_set": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arn": "string",
              "format_version": "string",
              "namespace": "string",
              "permission_policy": "string",
              "status": "string"
            }
          ]
        ]
      },
      "row_level_permission_tag_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "status": "string",
              "tag_rules": [
                "list",
                [
                  "object",
                  {
                    "column_name": "string",
                    "match_all_value": "string",
                    "tag_key": "string",
                    "tag_multi_value_delimiter": "string"
                  }
                ]
              ]
            }
          ]
        ]
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
    "block_types": {
      "column_level_permission_rules": {
        "block": {
          "attributes": {
            "column_names": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "principals": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsQuicksightDataSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsQuicksightDataSet), &result)
	return &result
}
