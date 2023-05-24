package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNetworkfirewallFirewall = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delete_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "key_id": "string",
              "type": "string"
            }
          ]
        ]
      },
      "firewall_policy_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_policy_change_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "firewall_status": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "capacity_usage_summary": [
                "set",
                [
                  "object",
                  {
                    "cidrs": [
                      "set",
                      [
                        "object",
                        {
                          "available_cidr_count": "number",
                          "ip_set_references": [
                            "set",
                            [
                              "object",
                              {
                                "resolved_cidr_count": "number"
                              }
                            ]
                          ],
                          "utilized_cidr_count": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "configuration_sync_state_summary": "string",
              "status": "string",
              "sync_states": [
                "set",
                [
                  "object",
                  {
                    "attachment": [
                      "list",
                      [
                        "object",
                        {
                          "endpoint_id": "string",
                          "status": "string",
                          "subnet_id": "string"
                        }
                      ]
                    ],
                    "availability_zone": "string"
                  }
                ]
              ]
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subnet_change_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "subnet_mapping": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "subnet_id": "string"
            }
          ]
        ]
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "update_token": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNetworkfirewallFirewallSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNetworkfirewallFirewall), &result)
	return &result
}
