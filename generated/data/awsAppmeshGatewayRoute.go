package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppmeshGatewayRoute = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_date": {
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
      "last_updated_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "mesh_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mesh_owner": {
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
      "resource_owner": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "spec": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "grpc_route": [
                "list",
                [
                  "object",
                  {
                    "action": [
                      "list",
                      [
                        "object",
                        {
                          "target": [
                            "list",
                            [
                              "object",
                              {
                                "port": "number",
                                "virtual_service": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "virtual_service_name": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "match": [
                      "list",
                      [
                        "object",
                        {
                          "port": "number",
                          "service_name": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "http2_route": [
                "list",
                [
                  "object",
                  {
                    "action": [
                      "list",
                      [
                        "object",
                        {
                          "rewrite": [
                            "list",
                            [
                              "object",
                              {
                                "hostname": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "default_target_hostname": "string"
                                    }
                                  ]
                                ],
                                "path": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "exact": "string"
                                    }
                                  ]
                                ],
                                "prefix": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "default_prefix": "string",
                                      "value": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "target": [
                            "list",
                            [
                              "object",
                              {
                                "port": "number",
                                "virtual_service": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "virtual_service_name": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "match": [
                      "list",
                      [
                        "object",
                        {
                          "header": [
                            "set",
                            [
                              "object",
                              {
                                "invert": "bool",
                                "match": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "exact": "string",
                                      "prefix": "string",
                                      "range": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "end": "number",
                                            "start": "number"
                                          }
                                        ]
                                      ],
                                      "regex": "string",
                                      "suffix": "string"
                                    }
                                  ]
                                ],
                                "name": "string"
                              }
                            ]
                          ],
                          "hostname": [
                            "list",
                            [
                              "object",
                              {
                                "exact": "string",
                                "suffix": "string"
                              }
                            ]
                          ],
                          "path": [
                            "list",
                            [
                              "object",
                              {
                                "exact": "string",
                                "regex": "string"
                              }
                            ]
                          ],
                          "port": "number",
                          "prefix": "string",
                          "query_parameter": [
                            "set",
                            [
                              "object",
                              {
                                "match": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "exact": "string"
                                    }
                                  ]
                                ],
                                "name": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "http_route": [
                "list",
                [
                  "object",
                  {
                    "action": [
                      "list",
                      [
                        "object",
                        {
                          "rewrite": [
                            "list",
                            [
                              "object",
                              {
                                "hostname": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "default_target_hostname": "string"
                                    }
                                  ]
                                ],
                                "path": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "exact": "string"
                                    }
                                  ]
                                ],
                                "prefix": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "default_prefix": "string",
                                      "value": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "target": [
                            "list",
                            [
                              "object",
                              {
                                "port": "number",
                                "virtual_service": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "virtual_service_name": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "match": [
                      "list",
                      [
                        "object",
                        {
                          "header": [
                            "set",
                            [
                              "object",
                              {
                                "invert": "bool",
                                "match": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "exact": "string",
                                      "prefix": "string",
                                      "range": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "end": "number",
                                            "start": "number"
                                          }
                                        ]
                                      ],
                                      "regex": "string",
                                      "suffix": "string"
                                    }
                                  ]
                                ],
                                "name": "string"
                              }
                            ]
                          ],
                          "hostname": [
                            "list",
                            [
                              "object",
                              {
                                "exact": "string",
                                "suffix": "string"
                              }
                            ]
                          ],
                          "path": [
                            "list",
                            [
                              "object",
                              {
                                "exact": "string",
                                "regex": "string"
                              }
                            ]
                          ],
                          "port": "number",
                          "prefix": "string",
                          "query_parameter": [
                            "set",
                            [
                              "object",
                              {
                                "match": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "exact": "string"
                                    }
                                  ]
                                ],
                                "name": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "priority": "number"
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
      "virtual_gateway_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAppmeshGatewayRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppmeshGatewayRoute), &result)
	return &result
}
