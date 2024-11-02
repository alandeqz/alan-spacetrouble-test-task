// Code generated by swaggo/swag at 2024-11-02 22:08:09.616977 +0500 +05 m=+1.081003501. DO NOT EDIT.

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
        "/v1/bookings": {
            "get": {
                "description": "Get all the Bookings for the flights by SpaceTrouble",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bookings"
                ],
                "summary": "Get the Bookings",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The limit of the response length",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "The offset of the response",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Booking"
                            }
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/errors.GenericErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.GenericErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create the Booking for the flight by SpaceTrouble",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bookings"
                ],
                "summary": "Create the Booking",
                "parameters": [
                    {
                        "description": "New Booking Request",
                        "name": "booking",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/booking.BookingRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Booking"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/errors.GenericErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.GenericErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/bookings/{id}": {
            "delete": {
                "description": "Delete the Booking for the flight by SpaceTrouble",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bookings"
                ],
                "summary": "Delete the Booking",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The ID of the Booking",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.GenericErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.GenericErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "booking.BookingRequest": {
            "type": "object",
            "properties": {
                "birthday": {
                    "type": "string",
                    "example": "1999-09-01T00:00:00Z"
                },
                "created_at": {
                    "type": "string",
                    "example": "2024-10-29T13:55:28.897Z"
                },
                "destination_id": {
                    "type": "string",
                    "enum": [
                        "Mars",
                        "Moon",
                        "Pluto",
                        "Asteroid Belt",
                        "Europa",
                        "Titan",
                        "Ganymede"
                    ]
                },
                "first_name": {
                    "type": "string",
                    "example": "John"
                },
                "gender": {
                    "$ref": "#/definitions/models.Gender"
                },
                "id": {
                    "type": "integer",
                    "example": 9
                },
                "last_name": {
                    "type": "string",
                    "example": "Doe"
                },
                "launch_date": {
                    "type": "string",
                    "example": "2024-12-01T15:00:00.000Z"
                },
                "launchpad_id": {
                    "type": "string",
                    "example": "1"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2024-10-30T12:13:37.374Z"
                }
            }
        },
        "errors.GenericErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.Booking": {
            "type": "object",
            "properties": {
                "birthday": {
                    "type": "string",
                    "example": "1999-09-01T00:00:00Z"
                },
                "created_at": {
                    "type": "string",
                    "example": "2024-10-29T13:55:28.897Z"
                },
                "destination_id": {
                    "type": "string",
                    "example": "2"
                },
                "first_name": {
                    "type": "string",
                    "example": "John"
                },
                "gender": {
                    "$ref": "#/definitions/models.Gender"
                },
                "id": {
                    "type": "integer",
                    "example": 9
                },
                "last_name": {
                    "type": "string",
                    "example": "Doe"
                },
                "launch_date": {
                    "type": "string",
                    "example": "2024-12-01T15:00:00.000Z"
                },
                "launchpad_id": {
                    "type": "string",
                    "example": "1"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2024-10-30T12:13:37.374Z"
                }
            }
        },
        "models.Gender": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3
            ],
            "x-enum-varnames": [
                "Unknown",
                "Male",
                "Female",
                "Other"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Alan Tabeo Test Task API",
	Description:      "This page contains the list of API specifications for the Tabeo test task.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
