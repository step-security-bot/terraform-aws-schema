package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsIamPrincipalPolicySimulation = `{
  "block": {
    "attributes": {
      "action_names": {
        "description": "One or more names of actions, like \"iam:CreateUser\", that should be included in the simulation.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "additional_policies_json": {
        "description": "Additional principal-based policies to use in the simulation.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "all_allowed": {
        "computed": true,
        "description": "A summary of the results attribute which is true if all of the results have decision \"allowed\", and false otherwise.",
        "description_kind": "plain",
        "type": "bool"
      },
      "caller_arn": {
        "description": "ARN of a user to use as the caller of the simulated requests. If not specified, defaults to the principal specified in policy_source_arn, if it is a user ARN.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Do not use",
        "description_kind": "plain",
        "type": "string"
      },
      "permissions_boundary_policies_json": {
        "description": "Additional permission boundary policies to use in the simulation.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "policy_source_arn": {
        "description": "ARN of the principal (e.g. user, role) whose existing configured access policies will be used as the basis for the simulation. If you specify a role ARN here, you can also set caller_arn to simulate a particular user acting with the given role.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_arns": {
        "description": "ARNs of specific resources to use as the targets of the specified actions during simulation. If not specified, the simulator assumes \"*\" which represents general access across all resources.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "resource_handling_option": {
        "description": "Specifies the type of simulation to run. Some API operations need a particular resource handling option in order to produce a correct reesult.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_owner_account_id": {
        "description": "An AWS account ID to use as the simulated owner for any resource whose ARN does not include a specific owner account ID. Defaults to the account given as part of caller_arn.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_policy_json": {
        "description": "A resource policy to associate with all of the target resources for simulation purposes. The policy simulator does not automatically retrieve resource-level policies, so if a resource policy is crucial to your test then you must specify here the same policy document associated with your target resource(s).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "results": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "action_name": "string",
              "allowed": "bool",
              "decision": "string",
              "decision_details": [
                "map",
                "string"
              ],
              "matched_statements": [
                "set",
                [
                  "object",
                  {
                    "source_policy_id": "string",
                    "source_policy_type": "string"
                  }
                ]
              ],
              "missing_context_keys": [
                "set",
                "string"
              ],
              "resource_arn": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "context": {
        "block": {
          "attributes": {
            "key": {
              "description": "The key name of the context entry, such as \"aws:CurrentTime\".",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description": "The type that the simulator should use to interpret the strings given in argument \"values\".",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description": "One or more values to assign to the context key, given as a string in a syntax appropriate for the selected value type.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description": "Each block specifies one item of additional context entry to include in the simulated requests. These are the additional properties used in the 'Condition' element of an IAM policy, and in dynamic value interpolations.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsIamPrincipalPolicySimulationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsIamPrincipalPolicySimulation), &result)
	return &result
}
