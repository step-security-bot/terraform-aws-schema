package resource

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
      }
    },
    "block_types": {
      "spec": {
        "block": {
          "block_types": {
            "backend_defaults": {
              "block": {
                "block_types": {
                  "client_policy": {
                    "block": {
                      "block_types": {
                        "tls": {
                          "block": {
                            "attributes": {
                              "enforce": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "ports": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "number"
                                ]
                              }
                            },
                            "block_types": {
                              "certificate": {
                                "block": {
                                  "block_types": {
                                    "file": {
                                      "block": {
                                        "attributes": {
                                          "certificate_chain": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "private_key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "sds": {
                                      "block": {
                                        "attributes": {
                                          "secret_name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "validation": {
                                "block": {
                                  "block_types": {
                                    "subject_alternative_names": {
                                      "block": {
                                        "block_types": {
                                          "match": {
                                            "block": {
                                              "attributes": {
                                                "exact": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": [
                                                    "set",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "min_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "trust": {
                                      "block": {
                                        "block_types": {
                                          "acm": {
                                            "block": {
                                              "attributes": {
                                                "certificate_authority_arns": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": [
                                                    "set",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "file": {
                                            "block": {
                                              "attributes": {
                                                "certificate_chain": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          },
                                          "sds": {
                                            "block": {
                                              "attributes": {
                                                "secret_name": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "max_items": 1,
                                            "nesting_mode": "list"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "listener": {
              "block": {
                "block_types": {
                  "connection_pool": {
                    "block": {
                      "block_types": {
                        "grpc": {
                          "block": {
                            "attributes": {
                              "max_requests": {
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
                        "http": {
                          "block": {
                            "attributes": {
                              "max_connections": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "max_pending_requests": {
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
                        "http2": {
                          "block": {
                            "attributes": {
                              "max_requests": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
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
                        "interval_millis": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "path": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "port": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "protocol": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "timeout_millis": {
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
                  "port_mapping": {
                    "block": {
                      "attributes": {
                        "port": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "protocol": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "tls": {
                    "block": {
                      "attributes": {
                        "mode": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "certificate": {
                          "block": {
                            "block_types": {
                              "acm": {
                                "block": {
                                  "attributes": {
                                    "certificate_arn": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "file": {
                                "block": {
                                  "attributes": {
                                    "certificate_chain": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "private_key": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "sds": {
                                "block": {
                                  "attributes": {
                                    "secret_name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "validation": {
                          "block": {
                            "block_types": {
                              "subject_alternative_names": {
                                "block": {
                                  "block_types": {
                                    "match": {
                                      "block": {
                                        "attributes": {
                                          "exact": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": [
                                              "set",
                                              "string"
                                            ]
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "min_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              },
                              "trust": {
                                "block": {
                                  "block_types": {
                                    "file": {
                                      "block": {
                                        "attributes": {
                                          "certificate_chain": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    },
                                    "sds": {
                                      "block": {
                                        "attributes": {
                                          "secret_name": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "max_items": 1,
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            },
            "logging": {
              "block": {
                "block_types": {
                  "access_log": {
                    "block": {
                      "block_types": {
                        "file": {
                          "block": {
                            "attributes": {
                              "path": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "format": {
                                "block": {
                                  "attributes": {
                                    "text": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "json": {
                                      "block": {
                                        "attributes": {
                                          "key": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "string"
                                          }
                                        },
                                        "description_kind": "plain"
                                      },
                                      "nesting_mode": "list"
                                    }
                                  },
                                  "description_kind": "plain"
                                },
                                "max_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
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
