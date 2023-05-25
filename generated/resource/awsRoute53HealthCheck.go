package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsRoute53HealthCheck = `{
  "block": {
    "attributes": {
      "child_health_threshold": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "child_healthchecks": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "cloudwatch_alarm_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cloudwatch_alarm_region": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_sni": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "failure_threshold": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "fqdn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "insufficient_data_health_status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "invert_healthcheck": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ip_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "measure_latency": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "reference_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "regions": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "request_interval": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "resource_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "search_string": {
        "description_kind": "plain",
        "optional": true,
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
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsRoute53HealthCheckSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsRoute53HealthCheck), &result)
	return &result
}
