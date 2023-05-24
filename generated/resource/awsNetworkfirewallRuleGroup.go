package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsNetworkfirewallRuleGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capacity": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "description": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rules": {
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
      "tags_all": {
        "computed": true,
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
      },
      "update_token": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "encryption_configuration": {
        "block": {
          "attributes": {
            "key_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
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
      "rule_group": {
        "block": {
          "block_types": {
            "reference_sets": {
              "block": {
                "block_types": {
                  "ip_set_references": {
                    "block": {
                      "attributes": {
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "ip_set_reference": {
                          "block": {
                            "attributes": {
                              "reference_arn": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 5,
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "rule_variables": {
              "block": {
                "block_types": {
                  "ip_sets": {
                    "block": {
                      "attributes": {
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "ip_set": {
                          "block": {
                            "attributes": {
                              "definition": {
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
                    "nesting_mode": "set"
                  },
                  "port_sets": {
                    "block": {
                      "attributes": {
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "port_set": {
                          "block": {
                            "attributes": {
                              "definition": {
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
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "rules_source": {
              "block": {
                "attributes": {
                  "rules_string": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "rules_source_list": {
                    "block": {
                      "attributes": {
                        "generated_rules_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "target_types": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "targets": {
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
                  "stateful_rule": {
                    "block": {
                      "attributes": {
                        "action": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "header": {
                          "block": {
                            "attributes": {
                              "destination": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "destination_port": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "direction": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "protocol": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "source": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "source_port": {
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
                        "rule_option": {
                          "block": {
                            "attributes": {
                              "keyword": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "settings": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              }
                            },
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "set"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "stateless_rules_and_custom_actions": {
                    "block": {
                      "block_types": {
                        "custom_action": {
                          "block": {
                            "attributes": {
                              "action_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "action_definition": {
                                "block": {
                                  "block_types": {
                                    "publish_metric_action": {
                                      "block": {
                                        "block_types": {
                                          "dimension": {
                                            "block": {
                                              "attributes": {
                                                "value": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "min_items": 1,
                                            "nesting_mode": "set"
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
                          "nesting_mode": "set"
                        },
                        "stateless_rule": {
                          "block": {
                            "attributes": {
                              "priority": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "block_types": {
                              "rule_definition": {
                                "block": {
                                  "attributes": {
                                    "actions": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    }
                                  },
                                  "block_types": {
                                    "match_attributes": {
                                      "block": {
                                        "attributes": {
                                          "protocols": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "set",
                                              "number"
                                            ]
                                          }
                                        },
                                        "block_types": {
                                          "destination": {
                                            "block": {
                                              "attributes": {
                                                "address_definition": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "nesting_mode": "set"
                                          },
                                          "destination_port": {
                                            "block": {
                                              "attributes": {
                                                "from_port": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                },
                                                "to_port": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "nesting_mode": "set"
                                          },
                                          "source": {
                                            "block": {
                                              "attributes": {
                                                "address_definition": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "nesting_mode": "set"
                                          },
                                          "source_port": {
                                            "block": {
                                              "attributes": {
                                                "from_port": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                },
                                                "to_port": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "nesting_mode": "set"
                                          },
                                          "tcp_flag": {
                                            "block": {
                                              "attributes": {
                                                "flags": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": [
                                                    "set",
                                                    "string"
                                                  ]
                                                },
                                                "masks": {
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": [
                                                    "set",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "description_kind": "plain"
                                            },
                                            "nesting_mode": "set"
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
                          "min_items": 1,
                          "nesting_mode": "set"
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
            "stateful_rule_options": {
              "block": {
                "attributes": {
                  "rule_order": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsNetworkfirewallRuleGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsNetworkfirewallRuleGroup), &result)
	return &result
}
