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
        "/admin/product/addcategory": {
            "post": {
                "security": [
                    {
                        "ApiKeyHeaderAuth": []
                    }
                ],
                "description": "Add category by admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product Management"
                ],
                "summary": "Add category",
                "parameters": [
                    {
                        "description": "Category object",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Category"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Added Category details",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/product/addproduct": {
            "post": {
                "security": [
                    {
                        "ApiKeyHeaderAuth": []
                    }
                ],
                "description": "Add product by admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product Management"
                ],
                "summary": "Add product",
                "parameters": [
                    {
                        "description": "Product object",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Added product details",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/product/deleteproduct": {
            "post": {
                "security": [
                    {
                        "ApiKeyHeaderAuth": []
                    }
                ],
                "description": "Delete product by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product Management"
                ],
                "summary": "Delete product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "product id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Array of deleted product details ",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Product"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Product"
                            }
                        }
                    }
                }
            }
        },
        "/admin/product/editproduct": {
            "post": {
                "security": [
                    {
                        "ApiKeyHeaderAuth": []
                    }
                ],
                "description": "Edit product by admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product Management"
                ],
                "summary": "Edit product",
                "parameters": [
                    {
                        "description": "Product object",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Edit product details",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/product/updatecategory": {
            "patch": {
                "security": [
                    {
                        "ApiKeyHeaderAuth": []
                    }
                ],
                "description": "delete Category by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product Management"
                ],
                "summary": "delete category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "category_id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "delete Category  ",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Category"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Category"
                            }
                        }
                    }
                }
            }
        },
        "/admin/users/block": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Edit block collumn of user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Management"
                ],
                "summary": "Block/Unblock-User",
                "parameters": [
                    {
                        "description": "blocked user id",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.userBlock"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Array of user details ",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Users"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Users"
                            }
                        }
                    }
                }
            }
        },
        "/admin/users/deleteuser": {
            "post": {
                "security": [
                    {
                        "ApiKeyHeaderAuth": []
                    }
                ],
                "description": "Delete user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Management"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User's id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Array of user details ",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/users/searchbyemail": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "ApiKeyHeaderAuth": []
                    }
                ],
                "description": "find user by email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Management"
                ],
                "summary": "Search user by email",
                "parameters": [
                    {
                        "description": "User's email address",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Array of user details ",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Users"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Users"
                            }
                        }
                    }
                }
            }
        },
        "/admin/users/userlist": {
            "post": {
                "security": [
                    {
                        "ApiKeyHeaderAuth": []
                    }
                ],
                "description": "Retrive and display user list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Management"
                ],
                "summary": "ListProducts",
                "responses": {
                    "200": {
                        "description": "Array of user details ",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserDetails"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserDetails"
                            }
                        }
                    }
                }
            }
        },
        "/cart/addtocart/{id}": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Add product to the cart using product id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Cart"
                ],
                "summary": "Add to Cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product-id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/user/otplogin": {
            "post": {
                "description": "verify Phone number using OTP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Auth"
                ],
                "summary": "SendOTP",
                "parameters": [
                    {
                        "description": "Phone number",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.OTPData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "phone number",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserSignInResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserSignInResponse"
                            }
                        }
                    }
                }
            }
        },
        "/user/signin": {
            "post": {
                "description": "Sign in a user and return user details and a token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Auth"
                ],
                "summary": "UserSignIN",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User details and role",
                        "schema": {
                            "$ref": "#/definitions/models.UserSignInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserSignInResponse"
                            }
                        }
                    }
                }
            }
        },
        "/user/verifyotp": {
            "post": {
                "description": "verify Phone number using OTP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Auth"
                ],
                "summary": "VerifyOTP",
                "parameters": [
                    {
                        "description": "Phone number and code",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.VerifyData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "phone number",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserSignInResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserSignInResponse"
                            }
                        }
                    }
                }
            }
        },
        "/users/signup": {
            "get": {
                "description": "Retrive UserDetails stored in DB and a auth token with success message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Auth"
                ],
                "summary": "UserSignUP",
                "parameters": [
                    {
                        "description": "User details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserDetails"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User details and token",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserDetails"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserSignInResponse"
                            }
                        }
                    }
                }
            }
        },
        "/users/viewproducts": {
            "get": {
                "description": "view products by a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "ViewProducts",
                "responses": {
                    "200": {
                        "description": "List of products",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Category": {
            "type": "object",
            "properties": {
                "category_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "domain.Product": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "colour": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "string"
                },
                "product_image": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "product_name": {
                    "type": "string"
                },
                "stock": {
                    "type": "string"
                }
            }
        },
        "domain.Users": {
            "type": "object",
            "properties": {
                "blocked": {
                    "type": "boolean"
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
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 8
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "boolean"
                }
            }
        },
        "handler.userBlock": {
            "type": "object",
            "properties": {
                "blocked": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.OTPData": {
            "type": "object",
            "required": [
                "phone"
            ],
            "properties": {
                "phone": {
                    "type": "string"
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "colour": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "string"
                },
                "product_image": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "product_name": {
                    "type": "string"
                },
                "stock": {
                    "type": "string"
                }
            }
        },
        "models.UserDetails": {
            "type": "object",
            "properties": {
                "confirmpassword": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "boolean",
                    "default": false
                }
            }
        },
        "models.UserLogin": {
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
        "models.UserSignInResponse": {
            "type": "object",
            "properties": {
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
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "boolean",
                    "default": false
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.VerifyData": {
            "type": "object",
            "required": [
                "code",
                "phone"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
