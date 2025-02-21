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
            "name": "Nome do Contato",
            "url": "http://www.example.com",
            "email": "example@example.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/clients": {
            "get": {
                "description": "Retorna todos os clientes cadastrados na base de dados",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Obtém a lista de todos os clientes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/clients.Client"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "Cria um novo cliente com validações de CPF, e-mail, telefone, nome e endereço",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Cria um novo cliente",
                "parameters": [
                    {
                        "description": "Cliente a ser criado",
                        "name": "Client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/clients.Client"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/clients.Client"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/clients/count": {
            "get": {
                "description": "Retorna a quantidade total de clientes cadastrados no sistema.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Obter número total de clientes",
                "responses": {
                    "200": {
                        "description": "Total de clientes"
                    },
                    "500": {
                        "description": "Erro ao obter a contagem de clientes"
                    }
                }
            }
        },
        "/clients/cpf/{cpf}": {
            "get": {
                "description": "Retorna os dados de um cliente com base no CPF informado.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Buscar cliente por CPF",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CPF do Cliente",
                        "name": "cpf",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/clients.Client"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Cliente não encontrado"
                    }
                }
            }
        },
        "/clients/name/{name}": {
            "get": {
                "description": "Retorna os dados dos clientes com base no nome informado.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Buscar clientes por nome",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Nome do Cliente",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/clients.Client"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Nenhum cliente encontrado"
                    }
                }
            }
        },
        "/clients/{id}": {
            "get": {
                "description": "Retorna um cliente específico através do seu ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Obtém um cliente pelo ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do cliente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/clients.Client"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "put": {
                "description": "Atualiza os dados de um cliente existente através do seu ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Atualiza as informações de um cliente",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do cliente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Cliente com dados atualizados",
                        "name": "Client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/clients.Client"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/clients.Client"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "Remove um cliente da base de dados pelo seu ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clients"
                ],
                "summary": "Deleta um cliente pelo ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do cliente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/deliveries": {
            "get": {
                "description": "Retorna todas as entregas cadastradas na base de dados",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deliveries"
                ],
                "summary": "Obtém a lista de todas as entregas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/deliveries.Delivery"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "Cria uma nova entrega com validações de peso e status de pedido",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deliveries"
                ],
                "summary": "Cria uma nova entrega",
                "parameters": [
                    {
                        "description": "Entrega a ser criada",
                        "name": "Delivery",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/deliveries.Delivery"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/deliveries.Delivery"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/deliveries/city/{city}": {
            "get": {
                "description": "Retorna todas as entregas associadas a uma cidade.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deliveries"
                ],
                "summary": "Buscar entregas por cidade",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Nome da Cidade",
                        "name": "city",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/deliveries.Delivery"
                            }
                        }
                    },
                    "400": {
                        "description": "Cidade inválida"
                    },
                    "404": {
                        "description": "Nenhuma entrega encontrada"
                    }
                }
            }
        },
        "/deliveries/client/cpf/{cpf}": {
            "get": {
                "description": "Retorna todas as entregas associadas a um CPF.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deliveries"
                ],
                "summary": "Buscar entregas por CPF",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CPF do Cliente",
                        "name": "cpf",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/deliveries.Delivery"
                            }
                        }
                    },
                    "400": {
                        "description": "CPF inválido"
                    },
                    "404": {
                        "description": "Nenhuma entrega encontrada"
                    }
                }
            }
        },
        "/deliveries/client/name/{name}": {
            "get": {
                "description": "Retorna todas as entregas associadas ao nome de um cliente.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deliveries"
                ],
                "summary": "Buscar entregas por nome do cliente",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Nome do Cliente",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/deliveries.Delivery"
                            }
                        }
                    },
                    "400": {
                        "description": "Nome inválido"
                    },
                    "404": {
                        "description": "Nenhuma entrega encontrada"
                    }
                }
            }
        },
        "/deliveries/{id}": {
            "get": {
                "description": "Retorna uma entrega específica através do seu ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deliveries"
                ],
                "summary": "Obtém uma entrega pelo ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID da entrega",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/deliveries.Delivery"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "put": {
                "description": "Atualiza os dados de uma entrega existente através do seu ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deliveries"
                ],
                "summary": "Atualiza as informações de uma entrega",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID da entrega",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Entrega com dados atualizados",
                        "name": "Delivery",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/deliveries.Delivery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/deliveries.Delivery"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "Remove uma entrega da base de dados pelo seu ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deliveries"
                ],
                "summary": "Deleta uma entrega pelo ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID da entrega",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/deliveries/{id}/status": {
            "patch": {
                "description": "Atualiza o status de uma entrega com base no ID da entrega.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deliveries"
                ],
                "summary": "Atualizar status do pedido",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID da Entrega",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Novo Status",
                        "name": "status",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/deliveries.Delivery"
                        }
                    },
                    "400": {
                        "description": "Requisição inválida"
                    },
                    "404": {
                        "description": "Entrega não encontrada"
                    }
                }
            }
        }
    },
    "definitions": {
        "clients.Client": {
            "description": "Dados da entrega",
            "type": "object",
            "required": [
                "birth_date",
                "cnpj",
                "cpf",
                "email",
                "name",
                "phone"
            ],
            "properties": {
                "birth_date": {
                    "type": "string"
                },
                "cnpj": {
                    "type": "string"
                },
                "cpf": {
                    "type": "string"
                },
                "deliveries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/deliveries.Delivery"
                    }
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
                }
            }
        },
        "deliveries.Delivery": {
            "description": "Dados da entrega",
            "type": "object",
            "properties": {
                "bairro": {
                    "type": "string"
                },
                "cidade": {
                    "type": "string"
                },
                "client_cpf": {
                    "type": "string"
                },
                "client_name": {
                    "type": "string"
                },
                "complemento": {
                    "type": "string"
                },
                "estado": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "latitude": {
                    "type": "number"
                },
                "logradouro": {
                    "type": "string"
                },
                "longitude": {
                    "type": "number"
                },
                "numero": {
                    "type": "string"
                },
                "order_status": {
                    "type": "string"
                },
                "pais": {
                    "type": "string"
                },
                "test_name": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "",
	Description:      "Esta é a documentação da API de entregas e clientes utilizando Gin e Swagger.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
