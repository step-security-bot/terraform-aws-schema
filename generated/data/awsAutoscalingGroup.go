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
      "desired_capacity_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled_metrics": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
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
      "max_instance_lifetime": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "mixed_instances_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "instances_distribution": [
                "list",
                [
                  "object",
                  {
                    "on_demand_allocation_strategy": "string",
                    "on_demand_base_capacity": "number",
                    "on_demand_percentage_above_base_capacity": "number",
                    "spot_allocation_strategy": "string",
                    "spot_instance_pools": "number",
                    "spot_max_price": "string"
                  }
                ]
              ],
              "launch_template": [
                "list",
                [
                  "object",
                  {
                    "launch_template_specification": [
                      "list",
                      [
                        "object",
                        {
                          "launch_template_id": "string",
                          "launch_template_name": "string",
                          "version": "string"
                        }
                      ]
                    ],
                    "override": [
                      "list",
                      [
                        "object",
                        {
                          "instance_requirements": [
                            "list",
                            [
                              "object",
                              {
                                "accelerator_count": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "max": "number",
                                      "min": "number"
                                    }
                                  ]
                                ],
                                "accelerator_manufacturers": [
                                  "set",
                                  "string"
                                ],
                                "accelerator_names": [
                                  "set",
                                  "string"
                                ],
                                "accelerator_total_memory_mib": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "max": "number",
                                      "min": "number"
                                    }
                                  ]
                                ],
                                "accelerator_types": [
                                  "set",
                                  "string"
                                ],
                                "allowed_instance_types": [
                                  "set",
                                  "string"
                                ],
                                "bare_metal": "string",
                                "baseline_ebs_bandwidth_mbps": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "max": "number",
                                      "min": "number"
                                    }
                                  ]
                                ],
                                "burstable_performance": "string",
                                "cpu_manufacturers": [
                                  "set",
                                  "string"
                                ],
                                "excluded_instance_types": [
                                  "set",
                                  "string"
                                ],
                                "instance_generations": [
                                  "set",
                                  "string"
                                ],
                                "local_storage": "string",
                                "local_storage_types": [
                                  "set",
                                  "string"
                                ],
                                "memory_gib_per_vcpu": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "max": "number",
                                      "min": "number"
                                    }
                                  ]
                                ],
                                "memory_mib": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "max": "number",
                                      "min": "number"
                                    }
                                  ]
                                ],
                                "network_bandwidth_gbps": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "max": "number",
                                      "min": "number"
                                    }
                                  ]
                                ],
                                "network_interface_count": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "max": "number",
                                      "min": "number"
                                    }
                                  ]
                                ],
                                "on_demand_max_price_percentage_over_lowest_price": "number",
                                "require_hibernate_support": "bool",
                                "spot_max_price_percentage_over_lowest_price": "number",
                                "total_local_storage_gb": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "max": "number",
                                      "min": "number"
                                    }
                                  ]
                                ],
                                "vcpu_count": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "max": "number",
                                      "min": "number"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "instance_type": "string",
                          "launch_template_specification": [
                            "list",
                            [
                              "object",
                              {
                                "launch_template_id": "string",
                                "launch_template_name": "string",
                                "version": "string"
                              }
                            ]
                          ],
                          "weighted_capacity": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
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
      "predicted_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "suspended_processes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "tag": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "key": "string",
              "propagate_at_launch": "bool",
              "value": "string"
            }
          ]
        ]
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
      "traffic_source": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "identifier": "string",
              "type": "string"
            }
          ]
        ]
      },
      "vpc_zone_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "warm_pool": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "instance_reuse_policy": [
                "list",
                [
                  "object",
                  {
                    "reuse_on_scale_in": "bool"
                  }
                ]
              ],
              "max_group_prepared_capacity": "number",
              "min_size": "number",
              "pool_state": "string"
            }
          ]
        ]
      },
      "warm_pool_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
