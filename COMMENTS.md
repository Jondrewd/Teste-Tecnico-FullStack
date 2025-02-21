# Documentação do Projeto

## Arquitetura e Decisões de Design

O projeto foi desenvolvido utilizando a arquitetura **MVC (Model-View-Controller)** com algumas adaptações para atender às necessidades de uma API RESTful. Abaixo estão as principais decisões de design:

1. **Separação de Responsabilidades**:
   - **Handlers**: Responsáveis por receber as requisições HTTP, validar os dados e enviar as respostas.
   - **Services**: Contêm a lógica de negócio e interagem com os repositórios para acessar os dados.
   - **Repositories**: Responsáveis pela comunicação com o banco de dados ou outras fontes de dados.
   - **Models**: Representam as entidades do domínio e são utilizados para mapear os dados.

2. **Uso do Gin Framework**:
   - O framework **Gin** foi escolhido por sua performance e simplicidade para criar APIs RESTful em Go.
   - Ele facilita o roteamento, a manipulação de middlewares e a serialização/deserialização de JSON.

3. **Testes Automatizados**:
   - Foram implementados testes unitários e de integração utilizando a biblioteca padrão `testing` do Go, juntamente com o pacote `testify` para asserções e mocks.
   - Os testes cobrem os principais cenários de sucesso e falha para garantir a robustez do código.

4. **Mocks para Testes**:
   - A biblioteca `testify/mock` foi utilizada para criar mocks dos serviços, permitindo testar os handlers de forma isolada.

5. **Validação de Dados**:
   - A validação dos dados de entrada é feita diretamente nos handlers antes de chamar os serviços, garantindo que apenas dados válidos sejam processados.

6. **Tratamento de Erros**:
   - Erros são tratados de forma centralizada, com respostas HTTP apropriadas (por exemplo, `400` para dados inválidos, `404` para recursos não encontrados e `500` para erros internos).

## Bibliotecas de Terceiros Utilizadas

- **Gin**: Framework web para criação de APIs RESTful em Go.
  - Repositório: [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- **Testify**: Biblioteca para testes, incluindo asserções e mocks.
  - Repositório: [https://github.com/stretchr/testify](https://github.com/stretchr/testify)
- **GORM**: ORM para interação com o banco de dados (se aplicável).
  - Repositório: [https://github.com/go-gorm/gorm](https://github.com/go-gorm/gorm)

## Requisitos Obrigatórios que Não Foram Entregues

Todos os requisitos obrigatórios foram entregues conforme especificado. Não há funcionalidades pendentes.