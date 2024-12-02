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
            "name": "GitHub",
            "url": "https://github.com/AmadoMuerte"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login/refresh": {
            "post": {
                "description": "This endpoint allows users to refresh their access token using a valid refresh token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Refresh Access token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {refresh_token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backend_internal_http-server_handlers_auth.tokenResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_AmadoMuerte_sysStats_internal_lib_response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_AmadoMuerte_sysStats_internal_lib_response.errorResponse"
                        }
                    }
                }
            }
        },
        "/login/sign-in": {
            "post": {
                "description": "This endpoint allows user to sign in using their email and passwd.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign In",
                "parameters": [
                    {
                        "description": "Credentials for signing in",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/backend_internal_http-server_handlers_auth.Credentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backend_internal_http-server_handlers_auth.tokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_AmadoMuerte_sysStats_internal_lib_response.errorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_AmadoMuerte_sysStats_internal_lib_response.errorResponse"
                        }
                    }
                }
            }
        },
        "/login/sign-up": {
            "post": {
                "description": "This endpoint allows user to sign up using their email and passwd.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign Up",
                "parameters": [
                    {
                        "description": "Credentials for signing up",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/backend_internal_http-server_handlers_auth.Credentials"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/backend_internal_http-server_handlers_auth.singUpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_AmadoMuerte_sysStats_internal_lib_response.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_AmadoMuerte_sysStats_internal_lib_response.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "backend_internal_http-server_handlers_auth.Credentials": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "backend_internal_http-server_handlers_auth.singUpResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "backend_internal_http-server_handlers_auth.tokenResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "backend_internal_lib_response.errorResponse": {
            "description": "This structure is used to send error information in JSON format.",
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
        "github_com_AmadoMuerte_sysStats_internal_lib_response.errorResponse": {
            "description": "This structure is used to send error information in JSON format.",
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "sysStats API",
	Description:      "API для sysStats",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
