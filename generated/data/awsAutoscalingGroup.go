package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAutoscalingGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "default_cooldown": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "desired_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "health_check_grace_period": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "health_check_type": {
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
      "launch_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "launch_template": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "name": "string",
              "version": "string"
            }
          ]
        ]
      },
      "load_balancers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "max_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "min_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "new_instances_protected_from_scale_in": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "placement_group": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_linked_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_group_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "termination_policies": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "vpc_zone_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAutoscalingGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAutoscalingGroup), &result)
	return &result
}
