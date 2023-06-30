package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsVpcSecurityGroupIngressRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cidr_ipv4": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cidr_ipv6": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "from_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ip_protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "prefix_list_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "referenced_security_group_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_group_rule_id": {
        "computed": true,
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
        "type": [
          "map",
          "string"
        ]
      },
      "to_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsVpcSecurityGroupIngressRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsVpcSecurityGroupIngressRule), &result)
	return &result
}
