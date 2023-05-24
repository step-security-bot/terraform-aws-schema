package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsLightsailDistribution = `{
  "block": {
    "attributes": {
      "alternative_domain_names": {
        "computed": true,
        "description": "The alternate domain names of the distribution.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the distribution.",
        "description_kind": "plain",
        "type": "string"
      },
      "bundle_id": {
        "description": "The bundle ID to use for the distribution.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate_name": {
        "description": "The name of the SSL/TLS certificate attached to the distribution, if any.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the distribution was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "The domain name of the distribution.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_address_type": {
        "description": "The IP address type of the distribution.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "is_enabled": {
        "description": "Indicates whether the distribution is enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "computed": true,
        "description": "An object that describes the location of the distribution, such as the AWS Region and Availability Zone.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "availability_zone": "string",
              "region_name": "string"
            }
          ]
        ]
      },
      "name": {
        "description": "The name of the distribution.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "origin_public_dns": {
        "computed": true,
        "description": "The public DNS of the origin.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "The Lightsail resource type (e.g., Distribution).",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the distribution.",
        "description_kind": "plain",
        "type": "string"
      },
      "support_code": {
        "computed": true,
        "description": "The support code. Include this code in your email to support when you have questions about your Lightsail distribution. This code enables our support team to look up your Lightsail information more easily.",
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
      }
    },
    "block_types": {
      "cache_behavior": {
        "block": {
          "attributes": {
            "behavior": {
              "description": "The cache behavior for the specified path.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description": "The path to a directory or file to cached, or not cache. Use an asterisk symbol to specify wildcard directories (path/to/assets/*), and file types (*.html, *jpg, *js). Directories and file paths are case-sensitive.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "An array of objects that describe the per-path cache behavior of the distribution.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "cache_behavior_settings": {
        "block": {
          "attributes": {
            "allowed_http_methods": {
              "description": "The HTTP methods that are processed and forwarded to the distribution's origin.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cached_http_methods": {
              "description": "The HTTP method responses that are cached by your distribution.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "default_ttl": {
              "description": "The default amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the content has been updated.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "maximum_ttl": {
              "description": "The maximum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "minimum_ttl": {
              "description": "The minimum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "forwarded_cookies": {
              "block": {
                "attributes": {
                  "cookies_allow_list": {
                    "description": "The specific cookies to forward to your distribution's origin.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "option": {
                    "description": "Specifies which cookies to forward to the distribution's origin for a cache behavior: all, none, or allow-list to forward only the cookies specified in the cookiesAllowList parameter.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "An object that describes the cookies that are forwarded to the origin. Your content is cached based on the cookies that are forwarded.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "forwarded_headers": {
              "block": {
                "attributes": {
                  "headers_allow_list": {
                    "description": "The specific headers to forward to your distribution's origin.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "option": {
                    "description": "The headers that you want your distribution to forward to your origin and base caching on.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "An object that describes the headers that are forwarded to the origin. Your content is cached based on the headers that are forwarded.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "forwarded_query_strings": {
              "block": {
                "attributes": {
                  "option": {
                    "description": "Indicates whether the distribution forwards and caches based on query strings.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "query_strings_allowed_list": {
                    "description": "The specific query strings that the distribution forwards to the origin.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description": "An object that describes the query strings that are forwarded to the origin. Your content is cached based on the query strings that are forwarded.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "An object that describes the cache behavior settings of the distribution.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "default_cache_behavior": {
        "block": {
          "attributes": {
            "behavior": {
              "description": "The cache behavior of the distribution.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "An object that describes the default cache behavior of the distribution.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "origin": {
        "block": {
          "attributes": {
            "name": {
              "description": "The name of the origin resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "protocol_policy": {
              "description": "The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region_name": {
              "description": "The AWS Region name of the origin resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource_type": {
              "computed": true,
              "description": "The resource type of the origin resource (e.g., Instance).",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "An object that describes the origin resource of the distribution, such as a Lightsail instance, bucket, or load balancer.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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
            },
            "update": {
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

func AwsLightsailDistributionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsLightsailDistribution), &result)
	return &result
}
