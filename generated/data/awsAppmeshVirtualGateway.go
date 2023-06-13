package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppmeshVirtualGateway = `{
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
              "backend_defaults": [
                "list",
                [
                  "object",
                  {
                    "client_policy": [
                      "list",
                      [
                        "object",
                        {
                          "tls": [
                            "list",
                            [
                              "object",
                              {
                                "certificate": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "file": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "certificate_chain": "string",
                                            "private_key": "string"
                                          }
                                        ]
                                      ],
                                      "sds": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "secret_name": "string"
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ],
                                "enforce": "bool",
                                "ports": [
                                  "set",
                                  "number"
                                ],
                                "validation": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "subject_alternative_names": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "match": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "exact": [
                                                    "set",
                                                    "string"
                                                  ]
                                                }
                                              ]
                                            ]
                                          }
                                        ]
                                      ],
                                      "trust": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "acm": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "certificate_authority_arns": [
                                                    "set",
                                                    "string"
                                                  ]
                                                }
                                              ]
                                            ],
                                            "file": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "certificate_chain": "string"
                                                }
                                              ]
                                            ],
                                            "sds": [
                                              "list",
                                              [
                                                "object",
                                                {
                                                  "secret_name": "string"
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
                        }
                      ]
                    ]
                  }
                ]
              ],
              "listener": [
                "list",
                [
                  "object",
                  {
                    "connection_pool": [
                      "list",
                      [
                        "object",
                        {
                          "grpc": [
                            "list",
                            [
                              "object",
                              {
                                "max_requests": "number"
                              }
                            ]
                          ],
                          "http": [
                            "list",
                            [
                              "object",
                              {
                                "max_connections": "number",
                                "max_pending_requests": "number"
                              }
                            ]
                          ],
                          "http2": [
                            "list",
                            [
                              "object",
                              {
                                "max_requests": "number"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "health_check": [
                      "list",
                      [
                        "object",
                        {
                          "healthy_threshold": "number",
                          "interval_millis": "number",
                          "path": "string",
                          "port": "number",
                          "protocol": "string",
                          "timeout_millis": "number",
                          "unhealthy_threshold": "number"
                        }
                      ]
                    ],
                    "port_mapping": [
                      "list",
                      [
                        "object",
                        {
                          "port": "number",
                          "protocol": "string"
                        }
                      ]
                    ],
                    "tls": [
                      "list",
                      [
                        "object",
                        {
                          "certificate": [
                            "list",
                            [
                              "object",
                              {
                                "acm": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "certificate_arn": "string"
                                    }
                                  ]
                                ],
                                "file": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "certificate_chain": "string",
                                      "private_key": "string"
                                    }
                                  ]
                                ],
                                "sds": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "secret_name": "string"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "mode": "string",
                          "validation": [
                            "list",
                            [
                              "object",
                              {
                                "subject_alternative_names": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "match": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "exact": [
                                              "set",
                                              "string"
                                            ]
                                          }
                                        ]
                                      ]
                                    }
                                  ]
                                ],
                                "trust": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "file": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "certificate_chain": "string"
                                          }
                                        ]
                                      ],
                                      "sds": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "secret_name": "string"
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
                  }
                ]
              ],
              "logging": [
                "list",
                [
                  "object",
                  {
                    "access_log": [
                      "list",
                      [
                        "object",
                        {
                          "file": [
                            "list",
                            [
                              "object",
                              {
                                "format": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "json": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "key": "string",
                                            "value": "string"
                                          }
                                        ]
                                      ],
                                      "text": "string"
                                    }
                                  ]
                                ],
                                "path": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsAppmeshVirtualGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppmeshVirtualGateway), &result)
	return &result
}
