package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsSecurityGroupRule = `{
  "block": {
    "attributes": {
      "cidr_blocks": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "from_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_cidr_blocks": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "prefix_list_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "self": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "source_security_group_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "to_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "type": {
        "description": "Type of rule, ingress (inbound) or egress (outbound).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 2
}`

func AwsSecurityGroupRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsSecurityGroupRule), &result)
	return &result
}
