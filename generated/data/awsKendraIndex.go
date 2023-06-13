package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsKendraIndex = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_units": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "query_capacity_units": "number",
              "storage_capacity_units": "number"
            }
          ]
        ]
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "document_metadata_configuration_updates": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "name": "string",
              "relevance": [
                "list",
                [
                  "object",
                  {
                    "duration": "string",
                    "freshness": "bool",
                    "importance": "number",
                    "rank_order": "string",
                    "values_importance_map": [
                      "map",
                      "number"
                    ]
                  }
                ]
              ],
              "search": [
                "list",
                [
                  "object",
                  {
                    "displayable": "bool",
                    "facetable": "bool",
                    "searchable": "bool",
                    "sortable": "bool"
                  }
                ]
              ],
              "type": "string"
            }
          ]
        ]
      },
      "edition": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "error_message": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "index_statistics": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "faq_statistics": [
                "list",
                [
                  "object",
                  {
                    "indexed_question_answers_count": "number"
                  }
                ]
              ],
              "text_document_statistics": [
                "list",
                [
                  "object",
                  {
                    "indexed_text_bytes": "number",
                    "indexed_text_documents_count": "number"
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
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "server_side_encryption_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kms_key_id": "string"
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_context_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_group_resolution_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "user_group_resolution_mode": "string"
            }
          ]
        ]
      },
      "user_token_configurations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "json_token_type_configuration": [
                "list",
                [
                  "object",
                  {
                    "group_attribute_field": "string",
                    "user_name_attribute_field": "string"
                  }
                ]
              ],
              "jwt_token_type_configuration": [
                "list",
                [
                  "object",
                  {
                    "claim_regex": "string",
                    "group_attribute_field": "string",
                    "issuer": "string",
                    "key_location": "string",
                    "secrets_manager_arn": "string",
                    "url": "string",
                    "user_name_attribute_field": "string"
                  }
                ]
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsKendraIndexSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsKendraIndex), &result)
	return &result
}
