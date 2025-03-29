// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cluster/info": {
            "get": {
                "tags": [
                    "cluster"
                ],
                "summary": "Get Cluster Info",
                "operationId": "\"GetClusterInfo\"",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/client_interaction.GetClusterInfoResponse"
                        }
                    }
                }
            }
        },
        "/cluster/log": {
            "get": {
                "tags": [
                    "cluster"
                ],
                "summary": "Get Log",
                "operationId": "\"GetLog\"",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/client_interaction.GetLogResponse"
                        }
                    }
                }
            }
        },
        "/key/{key}": {
            "get": {
                "tags": [
                    "key"
                ],
                "summary": "Get Key Value",
                "operationId": "\"GetKeyValue\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/client_interaction.GetKeyResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Sets value for the given key. If the old value already exists, it is replaced by a new one.",
                "tags": [
                    "key"
                ],
                "summary": "Set Key Value",
                "operationId": "\"SetKeyValue\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "key",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/client_interaction.SetKeyValueRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/client_interaction.CommandResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/client_interaction.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes value for the given key",
                "tags": [
                    "key"
                ],
                "summary": "Delete Key Value",
                "operationId": "\"DeleteKeyValue\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/client_interaction.CommandResponse"
                        }
                    }
                }
            },
            "patch": {
                "tags": [
                    "key"
                ],
                "summary": "Compare And Set Key Value",
                "operationId": "\"CompareAndSetKeyValue\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "key",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/client_interaction.CompareAndSetKeyValueRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/client_interaction.CommandResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/client_interaction.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "client_interaction.ClusterInfo": {
            "type": "object",
            "required": [
                "commitIndex",
                "currentTerm",
                "lastApplied",
                "matchIndex",
                "nextIndex"
            ],
            "properties": {
                "commitIndex": {
                    "type": "integer"
                },
                "currentTerm": {
                    "type": "integer"
                },
                "lastApplied": {
                    "type": "integer"
                },
                "matchIndex": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "nextIndex": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "client_interaction.CommandResponse": {
            "type": "object",
            "required": [
                "isLeader",
                "leaderId",
                "requestId"
            ],
            "properties": {
                "isLeader": {
                    "type": "boolean"
                },
                "leaderId": {
                    "type": "string"
                },
                "requestId": {
                    "type": "string",
                    "format": "uuid"
                }
            }
        },
        "client_interaction.CompareAndSetKeyValueRequest": {
            "type": "object",
            "properties": {
                "newValue": {},
                "oldValue": {}
            }
        },
        "client_interaction.ErrorResponse": {
            "type": "object",
            "required": [
                "error"
            ],
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "client_interaction.GetClusterInfoResponse": {
            "type": "object",
            "required": [
                "isLeader",
                "leaderId"
            ],
            "properties": {
                "info": {
                    "$ref": "#/definitions/client_interaction.ClusterInfo"
                },
                "isLeader": {
                    "type": "boolean"
                },
                "leaderId": {
                    "type": "string"
                }
            }
        },
        "client_interaction.GetKeyResponse": {
            "type": "object",
            "required": [
                "code",
                "isLeader",
                "leaderId",
                "value"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "enum": [
                        "success",
                        "not_found"
                    ]
                },
                "isLeader": {
                    "type": "boolean"
                },
                "leaderId": {
                    "type": "string"
                },
                "value": {}
            }
        },
        "client_interaction.GetLogResponse": {
            "type": "object",
            "required": [
                "entries",
                "isLeader",
                "leaderId"
            ],
            "properties": {
                "entries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/client_interaction.LogEntry"
                    }
                },
                "isLeader": {
                    "type": "boolean"
                },
                "leaderId": {
                    "type": "string"
                }
            }
        },
        "client_interaction.LogCommand": {
            "type": "object",
            "required": [
                "id",
                "key",
                "subKey",
                "type"
            ],
            "properties": {
                "id": {
                    "type": "string",
                    "format": "uuid"
                },
                "key": {
                    "type": "string"
                },
                "newValue": {},
                "oldValue": {},
                "subKey": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "set",
                        "compare_and_set",
                        "delete",
                        "add_element"
                    ]
                }
            }
        },
        "client_interaction.LogEntry": {
            "type": "object",
            "required": [
                "command",
                "term"
            ],
            "properties": {
                "command": {
                    "$ref": "#/definitions/client_interaction.LogCommand"
                },
                "term": {
                    "type": "integer"
                }
            }
        },
        "client_interaction.SetKeyValueRequest": {
            "type": "object",
            "properties": {
                "value": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
