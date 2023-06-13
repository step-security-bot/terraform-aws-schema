package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsApigatewayv2DomainName = `{
  "block": {
    "attributes": {
      "api_mapping_selection_expression": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
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
      "domain_name_configuration": {
        "block": {
          "attributes": {
            "certificate_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "endpoint_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "hosted_zone_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ownership_verification_certificate_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_policy": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_domain_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "mutual_tls_authentication": {
        "block": {
          "attributes": {
            "truststore_uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "truststore_version": {
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
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
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

func AwsApigatewayv2DomainNameSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsApigatewayv2DomainName), &result)
	return &result
}
