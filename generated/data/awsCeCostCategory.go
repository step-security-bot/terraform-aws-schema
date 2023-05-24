package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsCeCostCategory = `{
  "block": {
    "attributes": {
      "cost_category_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_value": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "effective_end": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "effective_start": {
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "rule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "inherited_value": [
                "list",
                [
                  "object",
                  {
                    "dimension_key": "string",
                    "dimension_name": "string"
                  }
                ]
              ],
              "rule": [
                "list",
                [
                  "object",
                  {
                    "and": [
                      "set",
                      [
                        "object",
                        {
                          "cost_category": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "match_options": [
                                  "set",
                                  "string"
                                ],
                                "values": [
                                  "set",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "dimension": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "match_options": [
                                  "set",
                                  "string"
                                ],
                                "values": [
                                  "set",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "tags": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "match_options": [
                                  "set",
                                  "string"
                                ],
                                "values": [
                                  "set",
                                  "string"
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "cost_category": [
                      "list",
                      [
                        "object",
                        {
                          "key": "string",
                          "match_options": [
                            "set",
                            "string"
                          ],
                          "values": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "dimension": [
                      "list",
                      [
                        "object",
                        {
                          "key": "string",
                          "match_options": [
                            "set",
                            "string"
                          ],
                          "values": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "not": [
                      "list",
                      [
                        "object",
                        {
                          "cost_category": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "match_options": [
                                  "set",
                                  "string"
                                ],
                                "values": [
                                  "set",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "dimension": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "match_options": [
                                  "set",
                                  "string"
                                ],
                                "values": [
                                  "set",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "tags": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "match_options": [
                                  "set",
                                  "string"
                                ],
                                "values": [
                                  "set",
                                  "string"
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "or": [
                      "set",
                      [
                        "object",
                        {
                          "cost_category": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "match_options": [
                                  "set",
                                  "string"
                                ],
                                "values": [
                                  "set",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "dimension": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "match_options": [
                                  "set",
                                  "string"
                                ],
                                "values": [
                                  "set",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "tags": [
                            "list",
                            [
                              "object",
                              {
                                "key": "string",
                                "match_options": [
                                  "set",
                                  "string"
                                ],
                                "values": [
                                  "set",
                                  "string"
                                ]
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "tags": [
                      "list",
                      [
                        "object",
                        {
                          "key": "string",
                          "match_options": [
                            "set",
                            "string"
                          ],
                          "values": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "type": "string",
              "value": "string"
            }
          ]
        ]
      },
      "rule_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "split_charge_rule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "method": "string",
              "parameter": [
                "set",
                [
                  "object",
                  {
                    "type": "string",
                    "values": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "source": "string",
              "targets": [
                "set",
                "string"
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

func AwsCeCostCategorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsCeCostCategory), &result)
	return &result
}
