package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAmplifyDomainAssociation = `{
  "block": {
    "attributes": {
      "app_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_verification_dns_record": {
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
      "wait_for_verification": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "sub_domain": {
        "block": {
          "attributes": {
            "branch_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "dns_record": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "prefix": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "verified": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAmplifyDomainAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAmplifyDomainAssociation), &result)
	return &result
}
