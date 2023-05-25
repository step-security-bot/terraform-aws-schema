package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCloudfrontCachePolicy = `{
  "block": {
    "attributes": {
      "comment": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_ttl": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "etag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_ttl": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "min_ttl": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameters_in_cache_key_and_forwarded_to_origin": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cookies_config": [
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
              ],
              "enable_accept_encoding_brotli": "bool",
              "enable_accept_encoding_gzip": "bool",
              "headers_config": [
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
              ],
              "query_strings_config": [
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
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsCloudfrontCachePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCloudfrontCachePolicy), &result)
	return &result
}
