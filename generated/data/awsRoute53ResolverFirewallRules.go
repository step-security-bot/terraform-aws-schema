package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRoute53ResolverFirewallRules = `{
  "block": {
    "attributes": {
      "action": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "firewall_rule_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "firewall_rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "action": "string",
              "block_override_dns_type": "string",
              "block_override_domain": "string",
              "block_override_ttl": "number",
              "block_response": "string",
              "creation_time": "string",
              "creator_request_id": "string",
              "firewall_domain_list_id": "string",
              "firewall_rule_group_id": "string",
              "modification_time": "string",
              "name": "string",
              "priority": "number"
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
      "priority": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRoute53ResolverFirewallRulesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRoute53ResolverFirewallRules), &result)
	return &result
}
