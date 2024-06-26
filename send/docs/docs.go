// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

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
        "/api/v1/task:send": {
            "post": {
                "description": "Send task to tasker",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Send task to tasker",
                "parameters": [
                    {
                        "description": "Send Task Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SendTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/restful.successResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.SendTaskResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/restful.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Customer": {
            "type": "object",
            "required": [
                "email",
                "name",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "huy@gmail.com"
                },
                "name": {
                    "type": "string",
                    "example": "huy tran"
                },
                "phone": {
                    "type": "string",
                    "example": "+84948447524"
                }
            }
        },
        "model.Location": {
            "type": "object",
            "required": [
                "address"
            ],
            "properties": {
                "address": {
                    "type": "string",
                    "example": "60 Bến Thành, Quận 1, TP.Hồ Chí Minh"
                }
            }
        },
        "model.SendTaskRequest": {
            "type": "object",
            "required": [
                "assigned_location",
                "customer",
                "id",
                "tasker",
                "total",
                "working_details"
            ],
            "properties": {
                "assigned_location": {
                    "$ref": "#/definitions/model.Location"
                },
                "customer": {
                    "$ref": "#/definitions/model.Customer"
                },
                "id": {
                    "type": "string",
                    "example": "c4824b96-e4ca-49cf-aaea-156604799612"
                },
                "note": {
                    "type": "string",
                    "example": "Be here ontime"
                },
                "tasker": {
                    "$ref": "#/definitions/model.Tasker"
                },
                "total": {
                    "type": "number",
                    "example": 400
                },
                "working_details": {
                    "$ref": "#/definitions/model.WorkingDetails"
                }
            }
        },
        "model.SendTaskResponse": {
            "type": "object",
            "properties": {
                "message_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.Tasker": {
            "type": "object",
            "required": [
                "email",
                "name",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "\u003center your email here\u003e"
                },
                "name": {
                    "type": "string",
                    "example": "huy tran"
                },
                "phone": {
                    "type": "string",
                    "example": "+84948447525"
                }
            }
        },
        "model.WorkingDetails": {
            "type": "object",
            "required": [
                "from_time",
                "house_type",
                "service_types",
                "to_time"
            ],
            "properties": {
                "from_time": {
                    "type": "string",
                    "example": "2024-04-07T09:00:00+07:00"
                },
                "house_type": {
                    "type": "string",
                    "example": "two_room"
                },
                "service_types": {
                    "type": "array",
                    "minItems": 1,
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "cleaning",
                        "childcare"
                    ]
                },
                "to_time": {
                    "type": "string",
                    "example": "2024-04-07T11:00:00+07:00"
                }
            }
        },
        "restful.errorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "restful.successResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8082",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Send API",
	Description:      "This is send apis",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
