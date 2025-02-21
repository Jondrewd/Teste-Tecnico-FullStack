package deliveries

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Handler é uma struct que manipula as requisições HTTP relacionadas a entregas.
// Ele contém uma instância de um serviço (Service) para realizar as operações de negócio.
type Handler struct {
	Service Service
}

// CreateDelivery é um handler HTTP para criar uma nova entrega.
// Ele valida os dados recebidos, cria a entrega no banco de dados e retorna uma resposta apropriada.
// @Summary Cria uma nova entrega
// @Description Cria uma nova entrega com validações de peso e status de pedido
// @Tags Deliveries
// @Accept json
// @Produce json
// @Param Delivery body Delivery true "Entrega a ser criada"
// @Success 201 {object} Delivery
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /deliveries [post]
func (h *Handler) CreateDelivery(c *gin.Context) {
	var delivery Delivery

	// Faz o bind dos dados JSON recebidos na requisição para a struct Delivery.
	// Se houver erro no bind (por exemplo, JSON inválido), retorna um erro 400 (Bad Request).
	if err := c.ShouldBindJSON(&delivery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Valida os dados da entrega usando a função validateDelivery.
	// Se a validação falhar, retorna um erro 400 (Bad Request).
	if err := validateDelivery(&delivery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chama o método CreateDelivery do serviço para criar a entrega no banco de dados.
	// Se houver erro na criação, retorna um erro 500 (Internal Server Error).
	createdDelivery, err := h.Service.CreateDelivery(&delivery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create delivery"})
		return
	}

	// Retorna a entrega criada com status 201 (Created).
	c.JSON(http.StatusCreated, createdDelivery)
}

// GetDeliveries é um handler HTTP para retornar todas as entregas cadastradas.
// @Summary Obtém a lista de todas as entregas
// @Description Retorna todas as entregas cadastradas na base de dados
// @Tags Deliveries
// @Accept json
// @Produce json
// @Success 200 {array} Delivery
// @Failure 500 "Internal Server Error"
// @Router /deliveries [get]
func (h *Handler) GetDeliveries(c *gin.Context) {
	// Chama o método GetDeliveries do serviço para obter a lista de entregas.
	deliveries, err := h.Service.GetDeliveries()
	if err != nil {
		// Se houver erro ao buscar as entregas, retorna um erro 500 (Internal Server Error).
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch deliveries"})
		return
	}

	// Retorna a lista de entregas com status 200 (OK).
	c.JSON(http.StatusOK, deliveries)
}

// GetDeliveryByID é um handler HTTP para retornar uma entrega específica pelo ID.
// @Summary Obtém uma entrega pelo ID
// @Description Retorna uma entrega específica através do seu ID
// @Tags Deliveries
// @Accept json
// @Produce json
// @Param id path int true "ID da entrega"
// @Success 200 {object} Delivery
// @Failure 400 "Bad Request"
// @Router /deliveries/{id} [get]
func (h *Handler) GetDeliveryByID(c *gin.Context) {
	// Obtém o ID da entrega da URL e converte para uint.
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		// Se o ID não for um número válido, retorna um erro 400 (Bad Request).
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Chama o método GetDeliveryByID do serviço para buscar a entrega pelo ID.
	delivery, err := h.Service.GetDeliveryByID(uint(id))
	if err != nil {
		// Se a entrega não for encontrada, retorna um erro 404 (Not Found).
		c.JSON(http.StatusNotFound, gin.H{"error": "Delivery not found"})
		return
	}

	// Retorna a entrega encontrada com status 200 (OK).
	c.JSON(http.StatusOK, delivery)
}

// UpdateDelivery é um handler HTTP para atualizar os dados de uma entrega existente.
// @Summary Atualiza as informações de uma entrega
// @Description Atualiza os dados de uma entrega existente através do seu ID
// @Tags Deliveries
// @Accept json
// @Produce json
// @Param id path int true "ID da entrega"
// @Param Delivery body Delivery true "Entrega com dados atualizados"
// @Success 200 {object} Delivery
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /deliveries/{id} [put]
func (h *Handler) UpdateDelivery(c *gin.Context) {
	// Obtém o ID da entrega da URL e converte para uint.
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		// Se o ID não for um número válido, retorna um erro 400 (Bad Request).
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Faz o bind dos dados JSON recebidos na requisição para a struct Delivery.
	var delivery Delivery
	if err := c.ShouldBindJSON(&delivery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Valida os dados da entrega.
	if err := validateDelivery(&delivery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chama o método UpdateDelivery do serviço para atualizar a entrega no banco de dados.
	updatedDelivery, err := h.Service.UpdateDelivery(uint(id), &delivery)
	if err != nil {
		// Se houver erro na atualização, retorna um erro 500 (Internal Server Error).
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update delivery"})
		return
	}

	// Retorna a entrega atualizada com status 200 (OK).
	c.JSON(http.StatusOK, updatedDelivery)
}

// DeleteDelivery é um handler HTTP para deletar uma entrega pelo ID.
// @Summary Deleta uma entrega pelo ID
// @Description Remove uma entrega da base de dados pelo seu ID
// @Tags Deliveries
// @Accept json
// @Produce json
// @Param id path int true "ID da entrega"
// @Success 204 {object} nil
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /deliveries/{id} [delete]
func (h *Handler) DeleteDelivery(c *gin.Context) {
	// Obtém o ID da entrega da URL e converte para uint.
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		// Se o ID não for um número válido, retorna um erro 400 (Bad Request).
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Chama o método DeleteDelivery do serviço para deletar a entrega do banco de dados.
	err = h.Service.DeleteDelivery(uint(id))
	if err != nil {
		// Se houver erro na exclusão, retorna um erro 500 (Internal Server Error).
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete delivery"})
		return
	}

	// Retorna status 204 (No Content) para indicar que a exclusão foi bem-sucedida.
	c.JSON(http.StatusNoContent, nil)
}

// validateDelivery realiza a validação dos campos da entrega.
// Ele usa o pacote validator para validar campos obrigatórios e regras personalizadas.
func validateDelivery(delivery *Delivery) error {
	validate := validator.New()

	// Valida os campos obrigatórios da struct Delivery.
	// Se algum campo obrigatório estiver faltando ou for inválido, retorna um erro.
	if err := validate.Struct(delivery); err != nil {
		return fmt.Errorf("validation failed: %s", err.Error())
	}

	// Valida o peso da entrega (não pode ser negativo ou zero).
	if delivery.Weight <= 0 {
		return fmt.Errorf("weight must be greater than zero")
	}

	// Verifica se o status da entrega é válido.
	if !isValidOrderStatus(delivery.OrderStatus) {
		return fmt.Errorf("invalid order status, must be one of: 'Pendente', 'Enviado', 'Entregue', 'Cancelado'")
	}

	// Se todas as validações passarem, retorna nil (sem erros).
	return nil
}

// GetDeliveriesByCPF é um handler HTTP para buscar todas as entregas associadas a um CPF.
// @Summary Buscar entregas por CPF
// @Description Retorna todas as entregas associadas a um CPF.
// @Tags Deliveries
// @Accept  json
// @Produce  json
// @Param cpf path string true "CPF do Cliente"
// @Success 200 {array} Delivery
// @Failure 400 "CPF inválido"
// @Failure 404 "Nenhuma entrega encontrada"
// @Router /deliveries/client/cpf/{cpf} [get]
func (h *Handler) GetDeliveriesByCPF(c *gin.Context) {
	// Obtém o CPF da URL.
	cpf := c.Param("cpf")

	// Chama o método GetDeliveriesByCPF do serviço para buscar as entregas pelo CPF.
	deliveries, err := h.Service.GetDeliveriesByCPF(cpf)
	if err != nil {
		// Se nenhuma entrega for encontrada, retorna um erro 404 (Not Found).
		c.JSON(http.StatusNotFound, gin.H{"error": "Nenhuma entrega encontrada"})
		return
	}

	// Retorna a lista de entregas com status 200 (OK).
	c.JSON(http.StatusOK, deliveries)
}

// GetDeliveriesByCity é um handler HTTP para buscar todas as entregas associadas a uma cidade.
// @Summary Buscar entregas por cidade
// @Description Retorna todas as entregas associadas a uma cidade.
// @Tags Deliveries
// @Accept  json
// @Produce  json
// @Param city path string true "Nome da Cidade"
// @Success 200 {array} Delivery
// @Failure 400 "Cidade inválida"
// @Failure 404 "Nenhuma entrega encontrada"
// @Router /deliveries/city/{city} [get]
func (h *Handler) GetDeliveriesByCity(c *gin.Context) {
    // Obtém o nome da cidade da URL.
    city := c.Param("city")

    // Chama o método GetDeliveriesByCity do serviço para buscar as entregas pela cidade.
    deliveries, err := h.Service.GetDeliveriesByCity(city)
    if err != nil {
        // Se nenhuma entrega for encontrada, retorna um erro 404 (Not Found).
        c.JSON(http.StatusNotFound, gin.H{"error": "Nenhuma entrega encontrada"})
        return
    }

    // Retorna a lista de entregas com status 200 (OK).
    c.JSON(http.StatusOK, deliveries)
}
// GetDeliveriesByClientName é um handler HTTP para buscar todas as entregas associadas a um cliente pelo nome.
// @Summary Buscar entregas por nome do cliente
// @Description Retorna todas as entregas associadas ao nome de um cliente.
// @Tags Deliveries
// @Accept  json
// @Produce  json
// @Param name path string true "Nome do Cliente"
// @Success 200 {array} Delivery
// @Failure 400 "Nome inválido"
// @Failure 404 "Nenhuma entrega encontrada"
// @Router /deliveries/client/name/{name} [get]
func (h *Handler) GetDeliveriesByClientName(c *gin.Context) {
	// Obtém o nome do cliente da URL
	name := c.Param("name")

	// Chama o método FindByClientName do serviço para buscar as entregas associadas ao nome do cliente
	deliveries, err := h.Service.GetDeliveriesByClientName(name)
	if err != nil {
		// Se nenhuma entrega for encontrada, retorna um erro 404 (Not Found)
		c.JSON(http.StatusNotFound, gin.H{"error": "Nenhuma entrega encontrada"})
		return
	}

	// Retorna a lista de entregas com status 200 (OK)
	c.JSON(http.StatusOK, deliveries)
}


// UpdateOrderStatus é um handler HTTP para atualizar o status de uma entrega.
// @Summary Atualizar status do pedido
// @Description Atualiza o status de uma entrega com base no ID da entrega.
// @Tags Deliveries
// @Accept  json
// @Produce  json
// @Param id path int true "ID da Entrega"
// @Param status body string true "Novo Status"
// @Success 200 {object} Delivery
// @Failure 400 "Requisição inválida"
// @Failure 404 "Entrega não encontrada"
// @Router /deliveries/{id}/status [patch]
func (h *Handler) UpdateOrderStatus(c *gin.Context) {
	// Obtém o ID da entrega da URL e converte para uint.
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Faz o bind dos dados JSON recebidos na requisição para uma struct.
	var request struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida"})
		return
	}

	// Chama o método UpdateOrderStatus do serviço para atualizar o status da entrega.
	err = h.Service.UpdateOrderStatus(uint(id), request.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entrega não encontrada"})
		return
	}

	// Retorna uma mensagem de sucesso com status 200 (OK).
	c.JSON(http.StatusOK, gin.H{"message": "Status atualizado com sucesso"})
}