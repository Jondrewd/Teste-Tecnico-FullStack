package deliveries_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"delivery-api/internal/deliveries" // Importação do pacote deliveries
)

// MockService simula o comportamento do Service para testes.
type MockService struct {
	mock.Mock
}

// CreateDelivery simula a criação de uma entrega.
func (m *MockService) CreateDelivery(delivery *deliveries.Delivery) (*deliveries.Delivery, error) {
	args := m.Called(delivery)
	return args.Get(0).(*deliveries.Delivery), args.Error(1)
}

// GetDeliveries simula a busca de todas as entregas.
func (m *MockService) GetDeliveries() ([]deliveries.Delivery, error) {
	args := m.Called()
	return args.Get(0).([]deliveries.Delivery), args.Error(1)
}

// GetDeliveryByID simula a busca de uma entrega por ID.
func (m *MockService) GetDeliveryByID(id uint) (*deliveries.Delivery, error) {
	args := m.Called(id)
	return args.Get(0).(*deliveries.Delivery), args.Error(1)
}

// UpdateDelivery simula a atualização de uma entrega.
func (m *MockService) UpdateDelivery(id uint, delivery *deliveries.Delivery) (*deliveries.Delivery, error) {
	args := m.Called(id, delivery)
	return args.Get(0).(*deliveries.Delivery), args.Error(1)
}

// DeleteDelivery simula a exclusão de uma entrega.
func (m *MockService) DeleteDelivery(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// GetDeliveriesByCPF simula a busca de entregas por CPF.
func (m *MockService) GetDeliveriesByCPF(cpf string) ([]deliveries.Delivery, error) {
	args := m.Called(cpf)
	return args.Get(0).([]deliveries.Delivery), args.Error(1)
}

// GetDeliveriesByClientName simula a busca de entregas por nome do cliente.
func (m *MockService) GetDeliveriesByClientName(name string) ([]deliveries.Delivery, error) {
	args := m.Called(name)
	return args.Get(0).([]deliveries.Delivery), args.Error(1)
}

// GetDeliveriesByCity simula a busca de entregas por cidade.
func (m *MockService) GetDeliveriesByCity(city string) ([]deliveries.Delivery, error) {
	args := m.Called(city)
	return args.Get(0).([]deliveries.Delivery), args.Error(1)
}

// UpdateOrderStatus simula a atualização do status de uma entrega.
func (m *MockService) UpdateOrderStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

// setupRouter inicializa o router do Gin com o handler de entregas.
func setupRouter(service deliveries.Service) *gin.Engine {
	handler := deliveries.Handler{Service: service}
	router := gin.Default()
	router.POST("/deliveries", handler.CreateDelivery)
	router.GET("/deliveries", handler.GetDeliveries)
	router.GET("/deliveries/:id", handler.GetDeliveryByID)
	router.PUT("/deliveries/:id", handler.UpdateDelivery)
	router.DELETE("/deliveries/:id", handler.DeleteDelivery)
	router.GET("/deliveries/client/cpf/:cpf", handler.GetDeliveriesByCPF)
	router.GET("/deliveries/client/name/:name", handler.GetDeliveriesByClientName)
	router.GET("/deliveries/city/:city", handler.GetDeliveriesByCity)
	router.PATCH("/deliveries/:id/status", handler.UpdateOrderStatus)
	return router
}

// TestCreateDelivery_Success testa a criação de uma entrega com sucesso.
func TestCreateDelivery_Success(t *testing.T) {
	mockService := new(MockService)
	router := setupRouter(mockService)

	// Define a entrega que será retornada pelo mock
	delivery := deliveries.Delivery{
		ClientCPF:   "123.456.789-00",
		ClientName:  "João Silva",
		TestName:    "Teste de Entrega",
		Weight:      10.5,
		Logradouro:  "Rua das Flores",
		Numero:      "123",
		Bairro:      "Centro",
		Complemento: "Apto 101",
		Cidade:      "São Paulo",
		Estado:      "SP",
		Pais:        "Brasil",
		Latitude:    -23.5505,
		Longitude:   -46.6333,
		OrderStatus: "Pending",
	}

	// Configura o mock para retornar a entrega quando CreateDelivery for chamado
	mockService.On("CreateDelivery", &delivery).Return(&delivery, nil)

	// Cria a requisição POST para criar a entrega
	body, _ := json.Marshal(delivery)
	req, _ := http.NewRequest("POST", "/deliveries", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Executa a requisição
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica se o status da resposta é 201 (Created)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Decodifica a resposta JSON e verifica se os dados da entrega estão corretos
	var response deliveries.Delivery
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "João Silva", response.ClientName)
	assert.Equal(t, "123.456.789-00", response.ClientCPF)
	assert.Equal(t, "Teste de Entrega", response.TestName)
	assert.Equal(t, 10.5, response.Weight)
	assert.Equal(t, "Rua das Flores", response.Logradouro)
	assert.Equal(t, "123", response.Numero)
	assert.Equal(t, "Centro", response.Bairro)
	assert.Equal(t, "Apto 101", response.Complemento)
	assert.Equal(t, "São Paulo", response.Cidade)
	assert.Equal(t, "SP", response.Estado)
	assert.Equal(t, "Brasil", response.Pais)
	assert.Equal(t, -23.5505, response.Latitude)
	assert.Equal(t, -46.6333, response.Longitude)
	assert.Equal(t, "Pending", response.OrderStatus)

	// Verifica se o método CreateDelivery foi chamado com a entrega correta
	mockService.AssertCalled(t, "CreateDelivery", &delivery)
}

// TestCreateDelivery_InvalidData testa a criação de uma entrega com dados inválidos.
func TestCreateDelivery_InvalidData(t *testing.T) {
	mockService := new(MockService)
	router := setupRouter(mockService)

	// Define uma entrega com dados inválidos (cidade vazia)
	delivery := deliveries.Delivery{
		ClientCPF:   "123.456.789-00",
		ClientName:  "João Silva",
		TestName:    "Teste de Entrega",
		Weight:      10.5,
		Logradouro:  "Rua das Flores",
		Numero:      "123",
		Bairro:      "Centro",
		Complemento: "Apto 101",
		Cidade:      "", // Cidade vazia (inválido)
		Estado:      "SP",
		Pais:        "Brasil",
		Latitude:    -23.5505,
		Longitude:   -46.6333,
		OrderStatus: "Pending",
	}

	// Cria a requisição POST para criar a entrega
	body, _ := json.Marshal(delivery)
	req, _ := http.NewRequest("POST", "/deliveries", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Executa a requisição
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica se o status da resposta é 400 (Bad Request)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestGetDeliveriesByCPF_Success testa a busca de entregas por CPF com sucesso.
func TestGetDeliveriesByCPF_Success(t *testing.T) {
	mockService := new(MockService)
	router := setupRouter(mockService)

	// Define a lista de entregas que será retornada pelo mock
	deliveriesList := []deliveries.Delivery{
		{
			ClientCPF:   "123.456.789-00",
			ClientName:  "João Silva",
			TestName:    "Teste de Entrega",
			Weight:      10.5,
			Logradouro:  "Rua das Flores",
			Numero:      "123",
			Bairro:      "Centro",
			Complemento: "Apto 101",
			Cidade:      "São Paulo",
			Estado:      "SP",
			Pais:        "Brasil",
			Latitude:    -23.5505,
			Longitude:   -46.6333,
			OrderStatus: "Pending",
		},
	}

	// Configura o mock para retornar a lista de entregas quando GetDeliveriesByCPF for chamado
	mockService.On("GetDeliveriesByCPF", "123.456.789-00").Return(deliveriesList, nil)

	// Cria a requisição GET para buscar entregas por CPF
	req, _ := http.NewRequest("GET", "/deliveries/client/cpf/123.456.789-00", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica se o status da resposta é 200 (OK)
	assert.Equal(t, http.StatusOK, w.Code)

	// Decodifica a resposta JSON e verifica se os dados das entregas estão corretos
	var response []deliveries.Delivery
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, "João Silva", response[0].ClientName)
	assert.Equal(t, "123.456.789-00", response[0].ClientCPF)
	assert.Equal(t, "Teste de Entrega", response[0].TestName)
	assert.Equal(t, 10.5, response[0].Weight)
	assert.Equal(t, "Rua das Flores", response[0].Logradouro)
	assert.Equal(t, "123", response[0].Numero)
	assert.Equal(t, "Centro", response[0].Bairro)
	assert.Equal(t, "Apto 101", response[0].Complemento)
	assert.Equal(t, "São Paulo", response[0].Cidade)
	assert.Equal(t, "SP", response[0].Estado)
	assert.Equal(t, "Brasil", response[0].Pais)
	assert.Equal(t, -23.5505, response[0].Latitude)
	assert.Equal(t, -46.6333, response[0].Longitude)
	assert.Equal(t, "Pending", response[0].OrderStatus)

	// Verifica se o método GetDeliveriesByCPF foi chamado com o CPF correto
	mockService.AssertCalled(t, "GetDeliveriesByCPF", "123.456.789-00")
}

// TestGetDeliveriesByClientName_Success testa a busca de entregas por nome do cliente com sucesso.
func TestGetDeliveriesByClientName_Success(t *testing.T) {
	mockService := new(MockService)
	router := setupRouter(mockService)

	// Define a lista de entregas que será retornada pelo mock
	deliveriesList := []deliveries.Delivery{
		{
			ClientCPF:   "123.456.789-00",
			ClientName:  "João Silva",
			TestName:    "Teste de Entrega",
			Weight:      10.5,
			Logradouro:  "Rua das Flores",
			Numero:      "123",
			Bairro:      "Centro",
			Complemento: "Apto 101",
			Cidade:      "São Paulo",
			Estado:      "SP",
			Pais:        "Brasil",
			Latitude:    -23.5505,
			Longitude:   -46.6333,
			OrderStatus: "Pending",
		},
	}

	// Configura o mock para retornar a lista de entregas quando GetDeliveriesByClientName for chamado
	mockService.On("GetDeliveriesByClientName", "João Silva").Return(deliveriesList, nil)

	// Cria a requisição GET para buscar entregas por nome do cliente
	req, _ := http.NewRequest("GET", "/deliveries/client/name/João Silva", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica se o status da resposta é 200 (OK)
	assert.Equal(t, http.StatusOK, w.Code)

	// Decodifica a resposta JSON e verifica se os dados das entregas estão corretos
	var response []deliveries.Delivery
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, "João Silva", response[0].ClientName)
	assert.Equal(t, "123.456.789-00", response[0].ClientCPF)
	assert.Equal(t, "Teste de Entrega", response[0].TestName)
	assert.Equal(t, 10.5, response[0].Weight)
	assert.Equal(t, "Rua das Flores", response[0].Logradouro)
	assert.Equal(t, "123", response[0].Numero)
	assert.Equal(t, "Centro", response[0].Bairro)
	assert.Equal(t, "Apto 101", response[0].Complemento)
	assert.Equal(t, "São Paulo", response[0].Cidade)
	assert.Equal(t, "SP", response[0].Estado)
	assert.Equal(t, "Brasil", response[0].Pais)
	assert.Equal(t, -23.5505, response[0].Latitude)
	assert.Equal(t, -46.6333, response[0].Longitude)
	assert.Equal(t, "Pending", response[0].OrderStatus)

	// Verifica se o método GetDeliveriesByClientName foi chamado com o nome correto
	mockService.AssertCalled(t, "GetDeliveriesByClientName", "João Silva")
}

// TestGetDeliveriesByCity_Success testa a busca de entregas por cidade com sucesso.
func TestGetDeliveriesByCity_Success(t *testing.T) {
	mockService := new(MockService)
	router := setupRouter(mockService)

	// Define a lista de entregas que será retornada pelo mock
	deliveriesList := []deliveries.Delivery{
		{
			ClientCPF:   "123.456.789-00",
			ClientName:  "João Silva",
			TestName:    "Teste de Entrega",
			Weight:      10.5,
			Logradouro:  "Rua das Flores",
			Numero:      "123",
			Bairro:      "Centro",
			Complemento: "Apto 101",
			Cidade:      "São Paulo",
			Estado:      "SP",
			Pais:        "Brasil",
			Latitude:    -23.5505,
			Longitude:   -46.6333,
			OrderStatus: "Pending",
		},
	}

	// Configura o mock para retornar a lista de entregas quando GetDeliveriesByCity for chamado
	mockService.On("GetDeliveriesByCity", "São Paulo").Return(deliveriesList, nil)

	// Cria a requisição GET para buscar entregas por cidade
	req, _ := http.NewRequest("GET", "/deliveries/city/São Paulo", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica se o status da resposta é 200 (OK)
	assert.Equal(t, http.StatusOK, w.Code)

	// Decodifica a resposta JSON e verifica se os dados das entregas estão corretos
	var response []deliveries.Delivery
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, "João Silva", response[0].ClientName)
	assert.Equal(t, "123.456.789-00", response[0].ClientCPF)
	assert.Equal(t, "Teste de Entrega", response[0].TestName)
	assert.Equal(t, 10.5, response[0].Weight)
	assert.Equal(t, "Rua das Flores", response[0].Logradouro)
	assert.Equal(t, "123", response[0].Numero)
	assert.Equal(t, "Centro", response[0].Bairro)
	assert.Equal(t, "Apto 101", response[0].Complemento)
	assert.Equal(t, "São Paulo", response[0].Cidade)
	assert.Equal(t, "SP", response[0].Estado)
	assert.Equal(t, "Brasil", response[0].Pais)
	assert.Equal(t, -23.5505, response[0].Latitude)
	assert.Equal(t, -46.6333, response[0].Longitude)
	assert.Equal(t, "Pending", response[0].OrderStatus)

	// Verifica se o método GetDeliveriesByCity foi chamado com a cidade correta
	mockService.AssertCalled(t, "GetDeliveriesByCity", "São Paulo")
}

// TestUpdateOrderStatus_Success testa a atualização do status de uma entrega com sucesso.
func TestUpdateOrderStatus_Success(t *testing.T) {
	mockService := new(MockService)
	router := setupRouter(mockService)

	// Configura o mock para retornar sucesso quando UpdateOrderStatus for chamado
	mockService.On("UpdateOrderStatus", uint(1), "Shipped").Return(nil)

	// Cria a requisição PATCH para atualizar o status da entrega
	body, _ := json.Marshal(map[string]string{"status": "Shipped"})
	req, _ := http.NewRequest("PATCH", "/deliveries/1/status", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Executa a requisição
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica se o status da resposta é 200 (OK)
	assert.Equal(t, http.StatusOK, w.Code)

	// Verifica se o método UpdateOrderStatus foi chamado com os parâmetros corretos
	mockService.AssertCalled(t, "UpdateOrderStatus", uint(1), "Shipped")
}