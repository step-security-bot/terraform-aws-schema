package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudfrontOriginRequestPolicy = `{
  "block": {
    "attributes": {
      "comment": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cookies_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cookie_behavior": "string",
              "cookies": [
                "list",
                [
                  "object",
                  {
                    "items": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "etag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "headers_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "header_behavior": "string",
              "headers": [
                "list",
                [
                  "object",
                  {
                    "items": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "query_strings_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "query_string_behavior": "string",
              "query_strings": [
                "list",
                [
                  "object",
                  {
                    "items": [
                      "set",
                      "string"
                    ]
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

func AwsCloudfrontOriginRequestPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudfrontOriginRequestPolicy), &result)
	return &result
}
