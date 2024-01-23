// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Tomáš Boďa",
            "url": "https://github.com/dvandyy",
            "email": "tomasboda.dev@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Retun a hello message if everything is ok",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Welcome"
                ],
                "summary": "Show a hello message",
                "responses": {
                    "200": {
                        "description": "Return 'Hello from gopi!'",
                        "schema": {
                            "$ref": "#/definitions/models.HelloResponse"
                        }
                    }
                }
            }
        },
        "/boards/new": {
            "post": {
                "description": "Create a new board in database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Boards"
                ],
                "summary": "Create new board",
                "parameters": [
                    {
                        "description": "Create board with Title and Description",
                        "name": "CreateBoardRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateBoardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CreateBoardResponse"
                        }
                    }
                }
            }
        },
        "/boards/{uid}": {
            "get": {
                "description": "Return board with unique id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Boards"
                ],
                "summary": "Get board with UUID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Board unique ID",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Board"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Create new user in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "Create user with Email and Password.",
                        "name": "RegisterRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RegisterResponse"
                        }
                    }
                }
            }
        },
        "/users/{uid}": {
            "get": {
                "description": "Return user with unique id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user with UUID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User unique ID",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Board": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2024-01-22 17:03:50.283466+00"
                },
                "description": {
                    "type": "string",
                    "example": "Board description"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "Title"
                },
                "unique_id": {
                    "type": "string",
                    "example": "926e7309-12e4-4c50-824c-33737fb45f8a"
                }
            }
        },
        "models.CreateBoardRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "My description"
                },
                "title": {
                    "type": "string",
                    "example": "My title"
                }
            }
        },
        "models.CreateBoardResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "models.HelloResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "models.RegisterRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@email.com"
                },
                "password": {
                    "type": "string",
                    "example": "Password123\u0026"
                }
            }
        },
        "models.RegisterResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2024-01-22 17:03:50.283466+00"
                },
                "email": {
                    "type": "string",
                    "example": "email@email.com"
                },
                "first_name": {
                    "type": "string",
                    "example": "FirstName"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "last_name": {
                    "type": "string",
                    "example": "LastName"
                },
                "password": {
                    "type": "string",
                    "example": "hash"
                },
                "role": {
                    "type": "string",
                    "example": "Tester"
                },
                "unique_id": {
                    "type": "string",
                    "example": "926e7309-12e4-4c50-824c-33737fb45f8a"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:4000",
	BasePath:         "/api/v1/",
	Schemes:          []string{},
	Title:            "Gopi API",
	Description:      "REST api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
