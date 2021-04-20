// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/spheres/all": {
            "get": {
                "description": "Get all spheres",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spheres"
                ],
                "summary": "Spheres",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.spheresResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/spheres/skills": {
            "post": {
                "description": "Get all skills from sphere",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spheres"
                ],
                "summary": "Spheres",
                "parameters": [
                    {
                        "description": "Spheres input info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.SpheresInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.skillsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/user/auth/refresh": {
            "post": {
                "description": "User refresh tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User auth"
                ],
                "summary": "User Refresh Tokens",
                "parameters": [
                    {
                        "description": "Refresh tokens info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserRefreshInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.tokensResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/user/auth/sign-in": {
            "post": {
                "description": "User sign-in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User auth"
                ],
                "summary": "User SignIn",
                "parameters": [
                    {
                        "description": "Sign-in info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserSignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.tokensResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/user/auth/sign-up": {
            "post": {
                "description": "User sign-up",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User auth"
                ],
                "summary": "User SignUp",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Image [jpeg/png]",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Look at the userStringTemplate or entity.UserSignUpInput in models",
                        "name": "user",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "description": "User sign-up template",
                        "name": "userStringTemplate",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/entity.UserSignUpInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.signUpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/user/existence": {
            "post": {
                "description": "User existence",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User existence",
                "parameters": [
                    {
                        "description": "User existence info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserExistenceInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.userExistenceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "get": {
                "security": [
                    {
                        "UserAuth": []
                    }
                ],
                "description": "user sign out from all devices",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User logout",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/user/sign-out": {
            "get": {
                "security": [
                    {
                        "UserAuth": []
                    }
                ],
                "description": "user sign out from current device",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User signOut",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Skill": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.Sphere": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.SpheresInput": {
            "type": "object",
            "required": [
                "spheres"
            ],
            "properties": {
                "spheres": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Sphere"
                    }
                }
            }
        },
        "entity.UserExistenceInput": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "entity.UserRefreshInput": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "entity.UserSignInInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "entity.UserSignUpInput": {
            "type": "object",
            "required": [
                "birthDate",
                "email",
                "firstName",
                "inSearch",
                "lastName",
                "password",
                "skills",
                "spheres"
            ],
            "properties": {
                "birthDate": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "inSearch": {
                    "type": "boolean"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "skills": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Skill"
                    }
                },
                "spheres": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Sphere"
                    }
                }
            }
        },
        "v1.response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.signUpResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "imageURL": {
                    "type": "string"
                }
            }
        },
        "v1.skillsResponse": {
            "type": "object",
            "properties": {
                "skills": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Skill"
                    }
                }
            }
        },
        "v1.spheresResponse": {
            "type": "object",
            "properties": {
                "spheres": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Sphere"
                    }
                }
            }
        },
        "v1.tokensResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "v1.userExistenceResponse": {
            "type": "object",
            "properties": {
                "Exists": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "UserAuth": {
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
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1/",
	Schemes:     []string{},
	Title:       "Network API",
	Description: "API Server for Network OPD Project",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
