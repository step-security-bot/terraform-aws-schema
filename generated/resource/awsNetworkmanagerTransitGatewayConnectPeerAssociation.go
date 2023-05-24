package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNetworkmanagerTransitGatewayConnectPeerAssociation = `{
  "block": {
    "attributes": {
      "device_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "global_network_id": {
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
      "link_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_gateway_connect_peer_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
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

func AwsNetworkmanagerTransitGatewayConnectPeerAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNetworkmanagerTransitGatewayConnectPeerAssociation), &result)
	return &result
}
