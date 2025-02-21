# Delivery API

## Descrição

A **Delivery API** é um sistema para gerenciar entregas e clientes. Ela permite o cadastro de clientes e entregas, e a consulta de entregas por CPF do cliente. A API foi construída em **Go** utilizando o framework **Gin** e **GORM** para facilitar a persistência no banco de dados. A documentação da API foi gerada automaticamente utilizando o **Swagger**.

## Funcionalidades

- Cadastro de clientes
- Cadastro de entregas associadas aos clientes
- Busca de entregas por CPF do cliente
- Busca de entregas associadas a um cliente por nome

## Tecnologias Utilizadas

- **Go**: Linguagem de programação principal
- **Gin**: Framework web para Go
- **GORM**: ORM para Go
- **MySQL**: Banco de dados relacional
- **Swagger**: Documentação da API
- **Logrus**: Biblioteca para logging avançado

## Endpoints

### /clients/{cpf} [GET]

#### Descrição:
Busca um cliente pelo CPF.

#### Parâmetros:
- `cpf` (string) - CPF do cliente

#### Resposta:
- **200 OK**: Retorna os dados do cliente
- **404 Not Found**: Cliente não encontrado

---

### /clients/name/{name} [GET]

#### Descrição:
Busca múltiplos clientes pelo nome (parcial), filtrando nomes que começam com o termo fornecido.

#### Parâmetros:
- `name` (string) - Nome do cliente (parcial)

#### Resposta:
- **200 OK**: Retorna a lista de clientes encontrados
- **404 Not Found**: Nenhum cliente encontrado

---

### /deliveries/{cpf} [GET]

#### Descrição:
Busca todas as entregas associadas a um CPF.

#### Parâmetros:
- `cpf` (string) - CPF do cliente

#### Resposta:
- **200 OK**: Retorna a lista de entregas
- **404 Not Found**: Nenhuma entrega encontrada

---

### Dependências

- **Gin** - Framework web para Go
- **GORM** - ORM para Go
- **MySQL Driver** - Driver MySQL para GORM
- **Swagger** - Para documentação automática da API
- **GoMock** - Framework de mocks para testes unitários
- **Logrus** - Biblioteca para logging avançado

## Testes Unitários

Os testes unitários são realizados com **GoMock**, e são responsáveis por validar as interações de cada camada da aplicação, incluindo o banco de dados e a camada de serviços.

## Como Rodar o Projeto

1. Clone o repositório:
    ```bash
    git clone https://github.com/seu_usuario/delivery-api.git
    ```

2. Entre no diretório do projeto:
    ```bash
    cd delivery-api
    ```

3. Instale as dependências:
    ```bash
    go mod tidy
    ```

4. Rodar a aplicação:
    ```bash
    go run main.go
    ```

5. Acesse a documentação da API via Swagger em `http://localhost:8080/swagger/index.html`.

## Testes

Os testes unitários podem ser executados com o seguinte comando:

```bash
go test ./...
