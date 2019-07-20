// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-07-20 00:55:49.458497987 +0000 UTC m=+0.044632833

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample gin starter server.",
        "title": "Go Gin Starter API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "gyuhwan.a.kim@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/gghcode/go-gin-starterkit/blob/master/LICENSE"
        },
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "/api",
    "paths": {
        "/auth/token": {
            "post": {
                "description": "Get new access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth API"
                ],
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/auth.CreateAccessTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/auth.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid credential",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/healthy": {
            "get": {
                "description": "Get server healthy",
                "tags": [
                    "App API"
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/todos": {
            "get": {
                "description": "Get all todos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo API"
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/todo.TodoResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create new todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo API"
                ],
                "parameters": [
                    {
                        "description": "todo payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.CreateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ok",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.TodoResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid todo payload",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get todo by todo id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo API"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.TodoResponse"
                        }
                    },
                    "404": {
                        "description": "Not found entity",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update todo by todo id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo API"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "todo payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.CreateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.TodoResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid todo payload",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not found entity",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Remove todo by todo id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo API"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.TodoResponse"
                        }
                    },
                    "404": {
                        "description": "Not found entity",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Create new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "parameters": [
                    {
                        "description": "user payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ok",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid user payload",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update new user by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid user payload",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not found entity",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Remove user by user id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not found entity",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/{username}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get user by username",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Name",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not found entity",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.CreateAccessTokenRequest": {
            "type": "object",
            "required": [
                "password",
                "user_name"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "\u003cpassword\u003e"
                },
                "user_name": {
                    "type": "string",
                    "example": "\u003cuser name\u003e"
                }
            }
        },
        "auth.TokenResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "common.APIError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "common.ErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/common.APIError"
                    }
                }
            }
        },
        "todo.CreateTodoRequest": {
            "type": "object",
            "required": [
                "contents",
                "title"
            ],
            "properties": {
                "contents": {
                    "type": "string",
                    "example": "\u003cnew contents\u003e"
                },
                "title": {
                    "type": "string",
                    "example": "\u003cnew title\u003e"
                }
            }
        },
        "todo.TodoResponse": {
            "type": "object",
            "properties": {
                "contents": {
                    "type": "string"
                },
                "create_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "user.CreateUserRequest": {
            "type": "object",
            "required": [
                "password",
                "user_name"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "\u003cnew password\u003e"
                },
                "user_name": {
                    "type": "string",
                    "example": "\u003cnew user name\u003e"
                }
            }
        },
        "user.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "user_name": {
                    "type": "string",
                    "example": "\u003cnew user name\u003e"
                }
            }
        },
        "user.UserResponse": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "user_name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
