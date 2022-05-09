// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/GJLANC65/gintest",
            "email": "[protected]"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/segments": {
            "get": {
                "produces": [
                    "application/json",
                    "application/xml"
                ],
                "summary": "get all items in the segment list",
                "operationId": "get-segments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Segment"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json",
                    "application/xml"
                ],
                "produces": [
                    "application/json",
                    "application/xml"
                ],
                "summary": "add a new item to the segment list",
                "operationId": "create-segment",
                "parameters": [
                    {
                        "description": "segment data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Segment"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Segment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gintest_controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/segments/{id}": {
            "get": {
                "produces": [
                    "application/json",
                    "application/xml"
                ],
                "summary": "get a segment item by ID",
                "operationId": "get-segment-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "segment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Segment"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gintest_controllers.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gintest_controllers.ErrorResponse": {
            "type": "object",
            "properties": {
                "err": {
                    "description": "Err error",
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "github.com_gjlanc65_gintest_controllers.ErrorResponse": {
            "type": "object",
            "properties": {
                "err": {
                    "description": "Err error",
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "models.Segment": {
            "description": "trip Segment information",
            "type": "object",
            "required": [
                "localDate"
            ],
            "properties": {
                "details": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "localDate": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "who": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Go + Gin Trip Segment API",
	Description:      "This is a sample server. You can visit the GitHub repository at https://github.com/GJLANC65/gintest",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
