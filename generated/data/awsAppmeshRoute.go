package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppmeshRoute = `{
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
                          "weighted_target": [
                            "set",
                            [
                              "object",
                              {
                                "port": "number",
                                "virtual_node": "string",
                                "weight": "number"
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
                          "metadata": [
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
                          "method_name": "string",
                          "port": "number",
                          "prefix": "string",
                          "service_name": "string"
                        }
                      ]
                    ],
                    "retry_policy": [
                      "list",
                      [
                        "object",
                        {
                          "grpc_retry_events": [
                            "set",
                            "string"
                          ],
                          "http_retry_events": [
                            "set",
                            "string"
                          ],
                          "max_retries": "number",
                          "per_retry_timeout": [
                            "list",
                            [
                              "object",
                              {
                                "unit": "string",
                                "value": "number"
                              }
                            ]
                          ],
                          "tcp_retry_events": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "timeout": [
                      "list",
                      [
                        "object",
                        {
                          "idle": [
                            "list",
                            [
                              "object",
                              {
                                "unit": "string",
                                "value": "number"
                              }
                            ]
                          ],
                          "per_request": [
                            "list",
                            [
                              "object",
                              {
                                "unit": "string",
                                "value": "number"
                              }
                            ]
                          ]
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
                          "weighted_target": [
                            "set",
                            [
                              "object",
                              {
                                "port": "number",
                                "virtual_node": "string",
                                "weight": "number"
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
                          "method": "string",
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
                          ],
                          "scheme": "string"
                        }
                      ]
                    ],
                    "retry_policy": [
                      "list",
                      [
                        "object",
                        {
                          "http_retry_events": [
                            "set",
                            "string"
                          ],
                          "max_retries": "number",
                          "per_retry_timeout": [
                            "list",
                            [
                              "object",
                              {
                                "unit": "string",
                                "value": "number"
                              }
                            ]
                          ],
                          "tcp_retry_events": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "timeout": [
                      "list",
                      [
                        "object",
                        {
                          "idle": [
                            "list",
                            [
                              "object",
                              {
                                "unit": "string",
                                "value": "number"
                              }
                            ]
                          ],
                          "per_request": [
                            "list",
                            [
                              "object",
                              {
                                "unit": "string",
                                "value": "number"
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
                          "weighted_target": [
                            "set",
                            [
                              "object",
                              {
                                "port": "number",
                                "virtual_node": "string",
                                "weight": "number"
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
                          "method": "string",
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
                          ],
                          "scheme": "string"
                        }
                      ]
                    ],
                    "retry_policy": [
                      "list",
                      [
                        "object",
                        {
                          "http_retry_events": [
                            "set",
                            "string"
                          ],
                          "max_retries": "number",
                          "per_retry_timeout": [
                            "list",
                            [
                              "object",
                              {
                                "unit": "string",
                                "value": "number"
                              }
                            ]
                          ],
                          "tcp_retry_events": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "timeout": [
                      "list",
                      [
                        "object",
                        {
                          "idle": [
                            "list",
                            [
                              "object",
                              {
                                "unit": "string",
                                "value": "number"
                              }
                            ]
                          ],
                          "per_request": [
                            "list",
                            [
                              "object",
                              {
                                "unit": "string",
                                "value": "number"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "priority": "number",
              "tcp_route": [
                "list",
                [
                  "object",
                  {
                    "action": [
                      "list",
                      [
                        "object",
                        {
                          "weighted_target": [
                            "set",
                            [
                              "object",
                              {
                                "port": "number",
                                "virtual_node": "string",
                                "weight": "number"
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
                          "port": "number"
                        }
                      ]
                    ],
                    "timeout": [
                      "list",
                      [
                        "object",
                        {
                          "idle": [
                            "list",
                            [
                              "object",
                              {
                                "unit": "string",
                                "value": "number"
                              }
                            ]
                          ]
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "virtual_router_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAppmeshRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppmeshRoute), &result)
	return &result
}
