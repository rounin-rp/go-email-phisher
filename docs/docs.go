// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/emails": {
            "get": {
                "description": "Get all email templates",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emails"
                ],
                "summary": "List all email templates",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Email"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "creates a new email template",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emails"
                ],
                "summary": "Creates a new email template",
                "parameters": [
                    {
                        "description": "Email data",
                        "name": "link",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateEmailTemplate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Email"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/send-emails": {
            "post": {
                "description": "sends the phising email provided by template id to a group of users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emails"
                ],
                "summary": "Sends the phising mail to a group of users",
                "parameters": [
                    {
                        "description": "Email data",
                        "name": "link",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.SendEmailToUsersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "402": {
                        "description": "Payment Required",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/user-link": {
            "post": {
                "description": "set user link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Set User Links",
                "parameters": [
                    {
                        "description": "Link data",
                        "name": "link",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateLinkRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Links"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/user-links": {
            "get": {
                "description": "Get all users links",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "List all users links",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Links"
                            }
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Set has_opened to true",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Set has_opened to true",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Link ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Links"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gin.H": {
            "type": "object",
            "additionalProperties": {}
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "handlers.CreateEmailTemplate": {
            "type": "object",
            "required": [
                "message",
                "subject"
            ],
            "properties": {
                "link": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                }
            }
        },
        "handlers.CreateLinkRequest": {
            "type": "object",
            "required": [
                "email",
                "user_id"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "handlers.SendEmailToUsersRequest": {
            "type": "object",
            "required": [
                "template_id",
                "user_emails"
            ],
            "properties": {
                "template_id": {
                    "type": "integer"
                },
                "user_emails": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.UserEmailMap"
                    }
                }
            }
        },
        "handlers.UserEmailMap": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.Email": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Links": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "has_opened": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "magicLink": {
                    "type": "string"
                },
                "times_clicked": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "",
	Description:      "This is a sample server for an email API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
