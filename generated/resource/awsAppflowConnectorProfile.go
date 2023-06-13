package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsAppflowConnectorProfile = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connector_label": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connector_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "credentials_arn": {
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
      "kms_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "connector_profile_config": {
        "block": {
          "block_types": {
            "connector_profile_credentials": {
              "block": {
                "block_types": {
                  "amplitude": {
                    "block": {
                      "attributes": {
                        "api_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_key": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "custom_connector": {
                    "block": {
                      "attributes": {
                        "authentication_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "api_key": {
                          "block": {
                            "attributes": {
                              "api_key": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "api_secret_key": {
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
                        "basic": {
                          "block": {
                            "attributes": {
                              "password": {
                                "description_kind": "plain",
                                "required": true,
                                "sensitive": true,
                                "type": "string"
                              },
                              "username": {
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
                        "custom": {
                          "block": {
                            "attributes": {
                              "credentials_map": {
                                "description_kind": "plain",
                                "optional": true,
                                "sensitive": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "custom_authentication_type": {
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
                        "oauth2": {
                          "block": {
                            "attributes": {
                              "access_token": {
                                "description_kind": "plain",
                                "optional": true,
                                "sensitive": true,
                                "type": "string"
                              },
                              "client_id": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "client_secret": {
                                "description_kind": "plain",
                                "optional": true,
                                "sensitive": true,
                                "type": "string"
                              },
                              "refresh_token": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "oauth_request": {
                                "block": {
                                  "attributes": {
                                    "auth_code": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "redirect_uri": {
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
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "datadog": {
                    "block": {
                      "attributes": {
                        "api_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "application_key": {
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
                  "dynatrace": {
                    "block": {
                      "attributes": {
                        "api_token": {
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
                  "google_analytics": {
                    "block": {
                      "attributes": {
                        "access_token": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "client_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "client_secret": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "refresh_token": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "oauth_request": {
                          "block": {
                            "attributes": {
                              "auth_code": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "redirect_uri": {
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
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "honeycode": {
                    "block": {
                      "attributes": {
                        "access_token": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "refresh_token": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "oauth_request": {
                          "block": {
                            "attributes": {
                              "auth_code": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "redirect_uri": {
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
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "infor_nexus": {
                    "block": {
                      "attributes": {
                        "access_key_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "datakey": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_access_key": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "user_id": {
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
                  "marketo": {
                    "block": {
                      "attributes": {
                        "access_token": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "client_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "client_secret": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "oauth_request": {
                          "block": {
                            "attributes": {
                              "auth_code": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "redirect_uri": {
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
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "redshift": {
                    "block": {
                      "attributes": {
                        "password": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "username": {
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
                  "salesforce": {
                    "block": {
                      "attributes": {
                        "access_token": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "client_credentials_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "refresh_token": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "oauth_request": {
                          "block": {
                            "attributes": {
                              "auth_code": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "redirect_uri": {
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
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "sapo_data": {
                    "block": {
                      "block_types": {
                        "basic_auth_credentials": {
                          "block": {
                            "attributes": {
                              "password": {
                                "description_kind": "plain",
                                "required": true,
                                "sensitive": true,
                                "type": "string"
                              },
                              "username": {
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
                        "oauth_credentials": {
                          "block": {
                            "attributes": {
                              "access_token": {
                                "description_kind": "plain",
                                "optional": true,
                                "sensitive": true,
                                "type": "string"
                              },
                              "client_id": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "client_secret": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "refresh_token": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "oauth_request": {
                                "block": {
                                  "attributes": {
                                    "auth_code": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "redirect_uri": {
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
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "service_now": {
                    "block": {
                      "attributes": {
                        "password": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "username": {
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
                  "singular": {
                    "block": {
                      "attributes": {
                        "api_key": {
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
                  "slack": {
                    "block": {
                      "attributes": {
                        "access_token": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "client_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "client_secret": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "oauth_request": {
                          "block": {
                            "attributes": {
                              "auth_code": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "redirect_uri": {
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
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "snowflake": {
                    "block": {
                      "attributes": {
                        "password": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "username": {
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
                  "trendmicro": {
                    "block": {
                      "attributes": {
                        "api_secret_key": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "veeva": {
                    "block": {
                      "attributes": {
                        "password": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "username": {
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
                  "zendesk": {
                    "block": {
                      "attributes": {
                        "access_token": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "client_id": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "client_secret": {
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "oauth_request": {
                          "block": {
                            "attributes": {
                              "auth_code": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "redirect_uri": {
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
            "connector_profile_properties": {
              "block": {
                "block_types": {
                  "amplitude": {
                    "block": {
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "custom_connector": {
                    "block": {
                      "attributes": {
                        "profile_properties": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "oauth2_properties": {
                          "block": {
                            "attributes": {
                              "oauth2_grant_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "token_url": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "token_url_custom_properties": {
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
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "datadog": {
                    "block": {
                      "attributes": {
                        "instance_url": {
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
                  "dynatrace": {
                    "block": {
                      "attributes": {
                        "instance_url": {
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
                  "google_analytics": {
                    "block": {
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "honeycode": {
                    "block": {
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "infor_nexus": {
                    "block": {
                      "attributes": {
                        "instance_url": {
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
                  "marketo": {
                    "block": {
                      "attributes": {
                        "instance_url": {
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
                  "redshift": {
                    "block": {
                      "attributes": {
                        "bucket_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "bucket_prefix": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cluster_identifier": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "data_api_role_arn": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "database_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "database_url": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
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
                  "salesforce": {
                    "block": {
                      "attributes": {
                        "instance_url": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "is_sandbox_environment": {
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
                  "sapo_data": {
                    "block": {
                      "attributes": {
                        "application_host_url": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "application_service_path": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "client_number": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "logon_language": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "port_number": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "private_link_service_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "oauth_properties": {
                          "block": {
                            "attributes": {
                              "auth_code_url": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "oauth_scopes": {
                                "description_kind": "plain",
                                "required": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "token_url": {
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
                  "service_now": {
                    "block": {
                      "attributes": {
                        "instance_url": {
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
                  "singular": {
                    "block": {
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "slack": {
                    "block": {
                      "attributes": {
                        "instance_url": {
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
                  "snowflake": {
                    "block": {
                      "attributes": {
                        "account_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "bucket_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "bucket_prefix": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "private_link_service_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "region": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "stage": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "warehouse": {
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
                  "trendmicro": {
                    "block": {
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "veeva": {
                    "block": {
                      "attributes": {
                        "instance_url": {
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
                  "zendesk": {
                    "block": {
                      "attributes": {
                        "instance_url": {
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
  "version": 0
}`

func AwsAppflowConnectorProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsAppflowConnectorProfile), &result)
	return &result
}
