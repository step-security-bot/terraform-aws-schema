package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsDlmLifecyclePolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "execution_role_arn": {
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
      "state": {
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
      }
    },
    "block_types": {
      "policy_details": {
        "block": {
          "attributes": {
            "policy_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_locations": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "resource_types": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "target_tags": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "block_types": {
            "action": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "cross_region_copy": {
                    "block": {
                      "attributes": {
                        "target": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "encryption_configuration": {
                          "block": {
                            "attributes": {
                              "cmk_arn": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "encrypted": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "retain_rule": {
                          "block": {
                            "attributes": {
                              "interval": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "interval_unit": {
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
                    "max_items": 3,
                    "min_items": 1,
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "event_source": {
              "block": {
                "attributes": {
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "parameters": {
                    "block": {
                      "attributes": {
                        "description_regex": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "event_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "snapshot_owner": {
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
            "parameters": {
              "block": {
                "attributes": {
                  "exclude_boot_volume": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "no_reboot": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "schedule": {
              "block": {
                "attributes": {
                  "copy_tags": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "tags_to_add": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "variable_tags": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "create_rule": {
                    "block": {
                      "attributes": {
                        "cron_expression": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "interval": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval_unit": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "location": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "times": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "cross_region_copy_rule": {
                    "block": {
                      "attributes": {
                        "cmk_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "copy_tags": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "encrypted": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "target": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "deprecate_rule": {
                          "block": {
                            "attributes": {
                              "interval": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "interval_unit": {
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
                        "retain_rule": {
                          "block": {
                            "attributes": {
                              "interval": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "interval_unit": {
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
                    "max_items": 3,
                    "nesting_mode": "set"
                  },
                  "deprecate_rule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval_unit": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "fast_restore_rule": {
                    "block": {
                      "attributes": {
                        "availability_zones": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "count": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval_unit": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "retain_rule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval_unit": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "share_rule": {
                    "block": {
                      "attributes": {
                        "target_accounts": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "unshare_interval": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "unshare_interval_unit": {
                          "description_kind": "plain",
                          "optional": true,
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
              "max_items": 4,
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

func AwsDlmLifecyclePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsDlmLifecyclePolicy), &result)
	return &result
}
