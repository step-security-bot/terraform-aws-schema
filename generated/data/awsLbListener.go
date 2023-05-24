package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLbListener = `{
  "block": {
    "attributes": {
      "alpn_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_action": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "authenticate_cognito": [
                "list",
                [
                  "object",
                  {
                    "authentication_request_extra_params": [
                      "map",
                      "string"
                    ],
                    "on_unauthenticated_request": "string",
                    "scope": "string",
                    "session_cookie_name": "string",
                    "session_timeout": "number",
                    "user_pool_arn": "string",
                    "user_pool_client_id": "string",
                    "user_pool_domain": "string"
                  }
                ]
              ],
              "authenticate_oidc": [
                "list",
                [
                  "object",
                  {
                    "authentication_request_extra_params": [
                      "map",
                      "string"
                    ],
                    "authorization_endpoint": "string",
                    "client_id": "string",
                    "client_secret": "string",
                    "issuer": "string",
                    "on_unauthenticated_request": "string",
                    "scope": "string",
                    "session_cookie_name": "string",
                    "session_timeout": "number",
                    "token_endpoint": "string",
                    "user_info_endpoint": "string"
                  }
                ]
              ],
              "fixed_response": [
                "list",
                [
                  "object",
                  {
                    "content_type": "string",
                    "message_body": "string",
                    "status_code": "string"
                  }
                ]
              ],
              "forward": [
                "list",
                [
                  "object",
                  {
                    "stickiness": [
                      "list",
                      [
                        "object",
                        {
                          "duration": "number",
                          "enabled": "bool"
                        }
                      ]
                    ],
                    "target_group": [
                      "set",
                      [
                        "object",
                        {
                          "arn": "string",
                          "weight": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "order": "number",
              "redirect": [
                "list",
                [
                  "object",
                  {
                    "host": "string",
                    "path": "string",
                    "port": "string",
                    "protocol": "string",
                    "query": "string",
                    "status_code": "string"
                  }
                ]
              ],
              "target_group_arn": "string",
              "type": "string"
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
      "load_balancer_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "protocol": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ssl_policy": {
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
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
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

func AwsLbListenerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLbListener), &result)
	return &result
}
