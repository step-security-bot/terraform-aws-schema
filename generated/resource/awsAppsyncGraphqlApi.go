package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppsyncGraphqlApi = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authentication_type": {
        "description_kind": "plain",
        "required": true,
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
      "schema": {
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
      "uris": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "visibility": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "xray_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "additional_authentication_provider": {
        "block": {
          "attributes": {
            "authentication_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "lambda_authorizer_config": {
              "block": {
                "attributes": {
                  "authorizer_result_ttl_in_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "authorizer_uri": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "identity_validation_expression": {
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
            "openid_connect_config": {
              "block": {
                "attributes": {
                  "auth_ttl": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "client_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "iat_ttl": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "issuer": {
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
            "user_pool_config": {
              "block": {
                "attributes": {
                  "app_id_client_regex": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "aws_region": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "user_pool_id": {
                    "description_kind": "plain",
                    "required": true,
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
        "nesting_mode": "list"
      },
      "lambda_authorizer_config": {
        "block": {
          "attributes": {
            "authorizer_result_ttl_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "authorizer_uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "identity_validation_expression": {
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
      "log_config": {
        "block": {
          "attributes": {
            "cloudwatch_logs_role_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "exclude_verbose_content": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "field_log_level": {
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
      "openid_connect_config": {
        "block": {
          "attributes": {
            "auth_ttl": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "client_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "iat_ttl": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "issuer": {
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
      "user_pool_config": {
        "block": {
          "attributes": {
            "app_id_client_regex": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "aws_region": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "default_action": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "user_pool_id": {
              "description_kind": "plain",
              "required": true,
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

func AwsAppsyncGraphqlApiSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppsyncGraphqlApi), &result)
	return &result
}
