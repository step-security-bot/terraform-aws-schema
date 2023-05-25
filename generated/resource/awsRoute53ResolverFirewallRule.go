package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRoute53ResolverFirewallRule = `{
  "block": {
    "attributes": {
      "action": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "block_override_dns_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "block_override_domain": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "block_override_ttl": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "block_response": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "firewall_domain_list_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "firewall_rule_group_id": {
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
      "priority": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRoute53ResolverFirewallRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRoute53ResolverFirewallRule), &result)
	return &result
}
