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
        "/api/v1/auth/login": {
            "post": {
                "description": "Login user with username and password",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Register \u0026 Login"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Username \u0026 Password",
                        "name": "Credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Credentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User logged in successfully",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "token": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid JSON",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "401": {
                        "description": "Invalid credentials",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "description": "Register a new user with username and password",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Register \u0026 Login"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "Username \u0026 Password",
                        "name": "Credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Credentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User registered successfully",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "message": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid JSON",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/tasks": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get tasks with pagination, sorting by status/created_at, and filtering by status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Get tasks with pagination, sorting, and filtering",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Items per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort by title/status/description/created_at",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort order: asc/desc",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by task status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/utils.Task"
                            }
                        }
                    },
                    "400": {
                        "description": "Error Message",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to fetch tasks:Error",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Create a new task",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Create a task",
                "parameters": [
                    {
                        "description": "Task Details",
                        "name": "TaskDetails",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.TaskDetails"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task created successfully",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "message": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Validation Error",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/tasks/mark-done": {
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Mark multiple tasks as done concurrently using Goroutines",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Mark tasks as done concurrently",
                "parameters": [
                    {
                        "description": "Task IDs to mark as done",
                        "name": "task_ids",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "message": {
                                    "type": "string"
                                },
                                "updated_tasks": {
                                    "type": "array",
                                    "items": {
                                        "type": "string"
                                    }
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid Task Id",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to Update task",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/tasks/{id}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get a task by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Get task by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Task"
                        }
                    },
                    "404": {
                        "description": "Task not found",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Update an existing task by ID",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Update a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated Task Details",
                        "name": "TaskDetails",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.TaskDetails"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task updated successfully",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "message": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid Task Id",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "404": {
                        "description": "Task not found",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to update task",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Delete a task by ID",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Delete a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task deleted successfully",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "message": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid Task Id",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "404": {
                        "description": "Task not found",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to delete task",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.Credentials": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "handlers.TaskDetails": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "utils.Task": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Go-Gin Task Manager API",
	Description:      "This API allow users to create, read, update, and delete tasks. Users can register, log in, and access their own tasks only when authenticated.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}