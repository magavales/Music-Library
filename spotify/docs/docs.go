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
        "/library": {
            "get": {
                "description": "get songs using the query's parameters for paginating and filtering",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "get songs"
                ],
                "summary": "Get songs",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The query's parameters for paginating and filtering",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "The query's parameters for paginating and filtering",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "The query's parameters for paginating and filtering",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The query's parameters for paginating and filtering",
                        "name": "group_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The query's parameters for paginating and filtering",
                        "name": "song_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The query's parameters for paginating and filtering",
                        "name": "release_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The query's parameters for paginating and filtering",
                        "name": "text",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The query's parameters for paginating and filtering",
                        "name": "link",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Song"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "create song",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "create"
                ],
                "summary": "Create song",
                "parameters": [
                    {
                        "description": "The song's info",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/library/:id": {
            "get": {
                "description": "get song by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "get song"
                ],
                "summary": "Get song",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The song's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Delete song",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delete"
                ],
                "summary": "Delete song",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The song's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "deleted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "patch": {
                "description": "Update song",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Update"
                ],
                "summary": "Update song",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The song's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The song's info",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "updated",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/library/:id/text": {
            "get": {
                "description": "get the song's text using the query's parameters for paginating and filtering",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "get the song's text"
                ],
                "summary": "Get the song's text",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The query's parameters for paginating and filtering",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Song": {
            "type": "object",
            "properties": {
                "group_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "song_name": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "todo API",
	Description:      "API Server for Music-Library Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
