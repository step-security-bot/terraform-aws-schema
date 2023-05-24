package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsElasticsearchDomain = `{
  "block": {
    "attributes": {
      "access_policies": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "advanced_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "advanced_security_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "internal_user_database_enabled": "bool"
            }
          ]
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_tune_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "desired_state": "string",
              "maintenance_schedule": [
                "set",
                [
                  "object",
                  {
                    "cron_expression_for_recurrence": "string",
                    "duration": [
                      "list",
                      [
                        "object",
                        {
                          "unit": "string",
                          "value": "number"
                        }
                      ]
                    ],
                    "start_at": "string"
                  }
                ]
              ],
              "rollback_on_disable": "string"
            }
          ]
        ]
      },
      "cluster_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cold_storage_options": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ],
              "dedicated_master_count": "number",
              "dedicated_master_enabled": "bool",
              "dedicated_master_type": "string",
              "instance_count": "number",
              "instance_type": "string",
              "warm_count": "number",
              "warm_enabled": "bool",
              "warm_type": "string",
              "zone_awareness_config": [
                "list",
                [
                  "object",
                  {
                    "availability_zone_count": "number"
                  }
                ]
              ],
              "zone_awareness_enabled": "bool"
            }
          ]
        ]
      },
      "cognito_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "identity_pool_id": "string",
              "role_arn": "string",
              "user_pool_id": "string"
            }
          ]
        ]
      },
      "created": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "deleted": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "domain_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ebs_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ebs_enabled": "bool",
              "iops": "number",
              "throughput": "number",
              "volume_size": "number",
              "volume_type": "string"
            }
          ]
        ]
      },
      "elasticsearch_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_at_rest": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "kms_key_id": "string"
            }
          ]
        ]
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kibana_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "log_publishing_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "cloudwatch_log_group_arn": "string",
              "enabled": "bool",
              "log_type": "string"
            }
          ]
        ]
      },
      "node_to_node_encryption": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool"
            }
          ]
        ]
      },
      "processing": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "snapshot_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "automated_snapshot_start_hour": "number"
            }
          ]
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "vpc_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "availability_zones": [
                "set",
                "string"
              ],
              "security_group_ids": [
                "set",
                "string"
              ],
              "subnet_ids": [
                "set",
                "string"
              ],
              "vpc_id": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsElasticsearchDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsElasticsearchDomain), &result)
	return &result
}
