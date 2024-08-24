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
            "name": "API SUPPORT",
            "url": "http://www.swagger.io/support",
            "email": "esezer@egetechno.com"
        },
        "license": {
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cash": {
            "post": {
                "description": "Create a cash entry with description, amount, type, and optionally link it to a customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cash"
                ],
                "summary": "Create a cash entry",
                "parameters": [
                    {
                        "description": "CashEntry",
                        "name": "cashentry",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CashEntry"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Cash entry created successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Shop or customer not found.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to create cash entry.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/employee/{shop_id}": {
            "get": {
                "description": "List employees in shop",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "List employees",
                "parameters": [
                    {
                        "description": "Employee",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Employee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Employee"
                        }
                    },
                    "400": {
                        "description": "Invalid id.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Employees not found.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Add an employee with role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Add an employee",
                "parameters": [
                    {
                        "description": "Employee",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Employee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Employee added successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Shop not found.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to add employee.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove an employee from shop",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Remove an employee",
                "parameters": [
                    {
                        "description": "Employee",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Employee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Employee removed successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "The employee is not an employee of this workplace.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Employee not found.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to remove employee.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/employee/{user_id}": {
            "put": {
                "description": "Update an employee with role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Update an employee",
                "parameters": [
                    {
                        "description": "Employee",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EmployeeRoleUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Employee role updated successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Employee not found.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to update employee.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/shop": {
            "post": {
                "description": "Create a shop with name and adress",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shop"
                ],
                "summary": "Create a shop",
                "parameters": [
                    {
                        "description": "Shop",
                        "name": "shop",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Shop"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Shop created successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to create shop.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/shop/{shop_id}": {
            "put": {
                "description": "Update shop's name and adress",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shop"
                ],
                "summary": "Update shop",
                "parameters": [
                    {
                        "description": "Shop",
                        "name": "shop",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Shop"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Shop updated successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to update shop.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a shop",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shop"
                ],
                "summary": "Delete a shop",
                "parameters": [
                    {
                        "description": "Shop",
                        "name": "shop",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Shop"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Shop deleted successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid id.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Shop not found.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to delete shop.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/transaction": {
            "post": {
                "description": "Create a new transaction.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Create a transaction",
                "parameters": [
                    {
                        "description": "Transaction",
                        "name": "Transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CustomerTransaction"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Transaction created successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to create transaction",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/transaction/{transaction_id}": {
            "get": {
                "description": "Get a transaction.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Get a transaction",
                "parameters": [
                    {
                        "description": "Transaction",
                        "name": "Transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CustomerTransaction"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.CustomerTransaction"
                        }
                    },
                    "400": {
                        "description": "Failed to create transaction",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Failed to get transaction.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/get-profile": {
            "get": {
                "description": "Pulls the user's own profile information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetProfile a user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Invalid input.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "User not found.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Register a new user with name, email, password, role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login a new user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "token: ",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Register a new user with name, email, password, role (e.g. owner, employee)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User registered successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/update-password": {
            "put": {
                "description": "Update the user's password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "UpdatePassword a user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PasswordUpdateReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Password updated successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "User not found.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/update-profile": {
            "put": {
                "description": "Update the user's profile.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "UpdateProfile a user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User updated successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "User not found.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/close-account": {
            "delete": {
                "description": "Close account",
                "tags": [
                    "User"
                ],
                "summary": "Close a user own account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Account closed successfully.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/{user_id}/delete": {
            "delete": {
                "description": "Delete a user by ID",
                "tags": [
                    "User"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "User deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CashEntry": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "createdAt": {
                    "type": "string"
                },
                "customer": {
                    "$ref": "#/definitions/models.Customer"
                },
                "customerID": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "entryType": {
                    "description": "e.g., income, expense",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "shop": {
                    "$ref": "#/definitions/models.Shop"
                },
                "shopID": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.Customer": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "shop": {
                    "$ref": "#/definitions/models.Shop"
                },
                "shopID": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.CustomerTransaction": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "createdAt": {
                    "type": "string"
                },
                "customer": {
                    "$ref": "#/definitions/models.Customer"
                },
                "customerID": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "shop": {
                    "$ref": "#/definitions/models.Shop"
                },
                "shopID": {
                    "type": "integer"
                },
                "transactionType": {
                    "description": "e.g., credit, debit",
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.Employee": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "role": {
                    "description": "e.g. employee, manager",
                    "type": "string"
                },
                "shop": {
                    "$ref": "#/definitions/models.Shop"
                },
                "shopID": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.EmployeeRoleUpdate": {
            "type": "object",
            "properties": {
                "role": {
                    "type": "string"
                }
            }
        },
        "models.LoginReq": {
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
        "models.PasswordUpdateReq": {
            "type": "object",
            "properties": {
                "new_password": {
                    "type": "string"
                },
                "old_password": {
                    "type": "string"
                }
            }
        },
        "models.Shop": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/models.User"
                },
                "ownerID": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "description": "e.g., admin, owner, employee etc.",
                    "type": "string"
                },
                "updatedAt": {
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
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Defterdar-go",
	Description:      "finance app for tradesman",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
