package clients

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Handler é uma struct que contém um serviço (Service) para manipular as operações relacionadas a clientes.
// Ele atua como um intermediário entre as requisições HTTP e a lógica de negócio.
type Handler struct {
	Service Service
}

// CreateClient é um handler HTTP para criar um novo cliente.
// Ele valida os dados recebidos, cria o cliente no banco de dados e retorna uma resposta apropriada.
// @Summary Cria um novo cliente
// @Description Cria um novo cliente com validações de CPF, e-mail, telefone, nome e endereço
// @Tags Clients
// @Accept json
// @Produce json
// @Param Client body Client true "Cliente a ser criado"
// @Success 201 {object} Client
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /clients [post]
func (h *Handler) CreateClient(c *gin.Context) {
	var client Client

	// Faz o bind dos dados JSON recebidos na requisição para a struct Client.
	// Se houver erro no bind (por exemplo, JSON inválido), retorna um erro 400 (Bad Request).
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// Valida os dados do cliente usando a função validateClient.
	// Se a validação falhar, retorna um erro 400 (Bad Request).
	if err := validateClient(&client); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// Chama o método CreateClient do serviço para criar o cliente no banco de dados.
	// Se houver erro na criação, retorna um erro 500 (Internal Server Error).
	createdClient, err := h.Service.CreateClient(&client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create client"})
		return
	}

	// Retorna o cliente criado com status 201 (Created).
	c.JSON(http.StatusCreated, createdClient)
}

// validateClient realiza a validação dos campos do cliente.
// Ele usa o pacote validator para validar campos obrigatórios e expressões regulares para formatos específicos.
func validateClient(client *Client) error {
	// Cria uma nova instância do validator.
	validate := validator.New()

	// Valida os campos obrigatórios da struct Client.
	// Se algum campo obrigatório estiver faltando ou for inválido, retorna um erro.
	if err := validate.Struct(client); err != nil {
		return fmt.Errorf("validation failed: %s", err.Error())
	}

	// Valida o formato do CPF usando uma expressão regular.
	if !isValidCPF(client.CPF) {
		return fmt.Errorf("invalid CPF")
	}

	// Valida o formato do CNPJ usando uma expressão regular.
	if !isValidCNPJ(client.CNPJ) {
		return fmt.Errorf("invalid CNPJ")
	}

	// Valida o formato do e-mail usando uma expressão regular.
	if !isValidEmail(client.Email) {
		return fmt.Errorf("invalid email format")
	}

	// Valida o formato do telefone (formato brasileiro) usando uma expressão regular.
	if !isValidPhone(client.Phone) {
		return fmt.Errorf("invalid phone format")
	}

	// Valida se o nome tem pelo menos 3 caracteres.
	if len(client.Name) < 3 {
		return fmt.Errorf("name must be at least 3 characters long")
	}

	// Valida a data de nascimento. A pessoa deve ter pelo menos 18 anos.
	if !isValidBirthdate(client.BirthDate) {
		return fmt.Errorf("invalid birthdate or client is underage")
	}
	// Se todas as validações passarem, retorna nil (sem erros).
	return nil
}

// isValidCPF verifica se o CPF está no formato correto usando uma expressão regular.
// O formato esperado é XXX.XXX.XXX-XX.
func isValidCPF(cpf string) bool {
	re := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}\-\d{2}$`)
	return re.MatchString(cpf)
}

// isValidEmail verifica se o e-mail está em um formato válido usando uma expressão regular.
// O formato esperado é algo como "exemplo@dominio.com".
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// isValidPhone verifica se o telefone está no formato brasileiro usando uma expressão regular.
// O formato esperado é (XX) XXXX-XXXX ou (XX) XXXXX-XXXX.
func isValidPhone(phone string) bool {
	re := regexp.MustCompile(`^\(?\d{2}\)?\s?\d{4,5}-\d{4}$`)
	return re.MatchString(phone)
}
// isValidCNPJ verifica se o CNPJ está no formato correto usando uma expressão regular.
// O formato esperado é XX.XXX.XXX/XXXX-XX.
func isValidCNPJ(cnpj string) bool {
	re := regexp.MustCompile(`^\d{2}\.\d{3}\.\d{3}\/\d{4}\-\d{2}$`)
	return re.MatchString(cnpj)
}

// isValidBirthdate verifica se a data de nascimento é válida e se o cliente tem pelo menos 18 anos.
func isValidBirthdate(birthdate time.Time) bool {
	// Verifica se o cliente tem pelo menos 18 anos
	age := time.Now().Year() - birthdate.Year()
	if age < 18 {
		return false
	}

	// Se a data de nascimento é válida e a idade é maior ou igual a 18 anos, retorna true.
	return true
}
// GetClients é um handler HTTP para retornar todos os clientes cadastrados.
// @Summary Obtém a lista de todos os clientes
// @Description Retorna todos os clientes cadastrados na base de dados
// @Tags Clients
// @Accept json
// @Produce json
// @Success 200 {array} Client
// @Failure 500 "Internal Server Error"
// @Router /clients [get]
func (h *Handler) GetClients(c *gin.Context) {
	// Chama o método GetClients do serviço para obter a lista de clientes.
	clients, err := h.Service.GetClients()
	if err != nil {
		// Se houver erro ao buscar os clientes, retorna um erro 500 (Internal Server Error).
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch clients"})
		return
	}

	// Retorna a lista de clientes com status 200 (OK).
	c.JSON(http.StatusOK, clients)
}

// GetClientByID é um handler HTTP para retornar um cliente específico pelo ID.
// @Summary Obtém um cliente pelo ID
// @Description Retorna um cliente específico através do seu ID
// @Tags Clients
// @Accept json
// @Produce json
// @Param id path int true "ID do cliente"
// @Success 200 {object} Client
// @Failure 400 "Bad Request"
// @Router /clients/{id} [get]
func (h *Handler) GetClientByID(c *gin.Context) {
	// Obtém o ID do cliente da URL e converte para inteiro.
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Se o ID não for um número válido, retorna um erro 400 (Bad Request).
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid client ID"})
		return
	}

	// Chama o método GetClientByID do serviço para buscar o cliente pelo ID.
	client, err := h.Service.GetClientByID(uint(id))
	if err != nil {
		// Se o cliente não for encontrado, retorna um erro 404 (Not Found).
		c.JSON(http.StatusNotFound, map[string]string{"error": "Client not found"})
		return
	}

	// Retorna o cliente encontrado com status 200 (OK).
	c.JSON(http.StatusOK, client)
}

// UpdateClient é um handler HTTP para atualizar os dados de um cliente existente.
// @Summary Atualiza as informações de um cliente
// @Description Atualiza os dados de um cliente existente através do seu ID
// @Tags Clients
// @Accept json
// @Produce json
// @Param id path int true "ID do cliente"
// @Param Client body Client true "Cliente com dados atualizados"
// @Success 200 {object} Client
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /clients/{id} [put]
func (h *Handler) UpdateClient(c *gin.Context) {
	// Obtém o ID do cliente da URL e converte para inteiro.
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Se o ID não for um número válido, retorna um erro 400 (Bad Request).
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid client ID"})
		return
	}

	// Faz o bind dos dados JSON recebidos na requisição para a struct Client.
	var client Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// Valida os dados do cliente.
	if err := validateClient(&client); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// Chama o método UpdateClient do serviço para atualizar o cliente no banco de dados.
	updatedClient, err := h.Service.UpdateClient(uint(id), &client)
	if err != nil {
		// Se houver erro na atualização, retorna um erro 500 (Internal Server Error).
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update client"})
		return
	}

	// Retorna o cliente atualizado com status 200 (OK).
	c.JSON(http.StatusOK, updatedClient)
}

// DeleteClient é um handler HTTP para deletar um cliente pelo ID.
// @Summary Deleta um cliente pelo ID
// @Description Remove um cliente da base de dados pelo seu ID
// @Tags Clients
// @Accept json
// @Produce json
// @Param id path int true "ID do cliente"
// @Success 204 {object} nil
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /clients/{id} [delete]
func (h *Handler) DeleteClient(c *gin.Context) {
	// Obtém o ID do cliente da URL e converte para inteiro.
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Se o ID não for um número válido, retorna um erro 400 (Bad Request).
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid client ID"})
		return
	}

	// Chama o método DeleteClient do serviço para deletar o cliente do banco de dados.
	err = h.Service.DeleteClient(uint(id))
	if err != nil {
		// Se houver erro na exclusão, retorna um erro 500 (Internal Server Error).
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete client"})
		return
	}

	// Retorna status 204 (No Content) para indicar que a exclusão foi bem-sucedida.
	c.JSON(http.StatusNoContent, nil)
}

// GetClientByCPF é um handler HTTP para buscar um cliente pelo CPF.
// @Summary Buscar cliente por CPF
// @Description Retorna os dados de um cliente com base no CPF informado.
// @Tags Clients
// @Accept  json
// @Produce  json
// @Param cpf path string true "CPF do Cliente"
// @Success 200 {object} Client
// @Failure 400 "Bad Request"
// @Failure 404 "Cliente não encontrado"
// @Router /clients/{cpf} [get]
func (h *Handler) GetClientByCPF(c *gin.Context) {
	// Obtém o CPF do cliente da URL.
	cpf := c.Param("cpf")

	// Chama o método GetClientByCPF do serviço para buscar o cliente pelo CPF.
	client, err := h.Service.GetClientByCPF(cpf)
	if err != nil {
		// Se o cliente não for encontrado, retorna um erro 404 (Not Found).
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	// Retorna o cliente encontrado com status 200 (OK).
	c.JSON(http.StatusOK, client)
}

// GetTotalClients é um handler HTTP para retornar o número total de clientes cadastrados.
// @Summary Obter número total de clientes
// @Description Retorna a quantidade total de clientes cadastrados no sistema.
// @Tags Clients
// @Accept  json
// @Produce  json
// @Success 200 "Total de clientes"
// @Failure 500 "Erro ao obter a contagem de clientes"
// @Router /clients/count [get]
func (h *Handler) GetTotalClients(c *gin.Context) {
	// Chama o método GetTotalClients do serviço para obter a contagem de clientes.
	count, err := h.Service.GetTotalClients()
	if err != nil {
		// Se houver erro ao obter a contagem, retorna um erro 500 (Internal Server Error).
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter a contagem de clientes"})
		return
	}

	// Retorna o total de clientes com status 200 (OK).
	c.JSON(http.StatusOK, gin.H{"total_clients": count})
}