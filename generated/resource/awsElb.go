package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsElb = `{
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
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "connection_draining": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "connection_draining_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cross_zone_load_balancing": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "desync_mitigation_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_name": {
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
      "idle_timeout": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "internal": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_prefix": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_groups": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "source_security_group": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_security_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnets": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
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
      "tags_all": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "access_logs": {
        "block": {
          "attributes": {
            "bucket": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "bucket_prefix": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "interval": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "health_check": {
        "block": {
          "attributes": {
            "healthy_threshold": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "interval": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "target": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "timeout": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "unhealthy_threshold": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "listener": {
        "block": {
          "attributes": {
            "instance_port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "instance_protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "lb_port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "lb_protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ssl_certificate_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsElbSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsElb), &result)
	return &result
}
