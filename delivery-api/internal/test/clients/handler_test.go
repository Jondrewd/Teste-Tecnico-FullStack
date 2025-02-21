package clients

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"delivery-api/internal/clients"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockService simula o comportamento do Service para testes.
// Ele implementa a interface Service e usa a biblioteca "testify/mock" para controlar as respostas.
type MockService struct {
	mock.Mock
}

// CreateClient simula a criação de um cliente.
func (m *MockService) CreateClient(client *clients.Client) (*clients.Client, error) {
	args := m.Called(client)
	return args.Get(0).(*clients.Client), args.Error(1)
}

// GetClients simula a busca de todos os clientes.
func (m *MockService) GetClients() ([]clients.Client, error) {
	args := m.Called()
	return args.Get(0).([]clients.Client), args.Error(1)
}

// GetClientByID simula a busca de um cliente por ID.
func (m *MockService) GetClientByID(id uint) (*clients.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*clients.Client), args.Error(1)
}

// UpdateClient simula a atualização de um cliente.
func (m *MockService) UpdateClient(id uint, client *clients.Client) (*clients.Client, error) {
	args := m.Called(id, client)
	return args.Get(0).(*clients.Client), args.Error(1)
}

// DeleteClient simula a exclusão de um cliente.
func (m *MockService) DeleteClient(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// GetClientByCPF simula a busca de um cliente por CPF.
func (m *MockService) GetClientByCPF(cpf string) (*clients.Client, error) {
	args := m.Called(cpf)
	return args.Get(0).(*clients.Client), args.Error(1)
}

// GetClientByName simula a busca de clientes por nome.
func (m *MockService) GetClientByName(name string) ([]clients.Client, error) {
	args := m.Called(name)
	return args.Get(0).([]clients.Client), args.Error(1)
}

// GetTotalClients simula a obtenção da contagem total de clientes.
func (m *MockService) GetTotalClients() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

// setupRouter inicializa o router do Gin com o handler de clientes.
// Ele configura todas as rotas necessárias para os testes.
func setupRouter(service clients.Service) *gin.Engine {
	handler := clients.Handler{Service: service}
	router := gin.Default()
	router.POST("/clients", handler.CreateClient)
	router.GET("/clients", handler.GetClients)
	router.GET("/clients/:id", handler.GetClientByID)
	router.PUT("/clients/:id", handler.UpdateClient)
	router.DELETE("/clients/:id", handler.DeleteClient)
	router.GET("/clients/cpf/:cpf", handler.GetClientByCPF)
	router.GET("/clients/name/:name", handler.GetClientsByName)
	router.GET("/clients/count", handler.GetTotalClients)
	return router
}

// TestGetClientByCPF_Success testa a busca de um cliente por CPF com sucesso.
func TestGetClientByCPF_Success(t *testing.T) {
	mockService := new(MockService)
	router := setupRouter(mockService)

	// Define o cliente que será retornado pelo mock
	client := clients.Client{
		ID:    1,
		Name:  "João Silva",
		CPF:   "123.456.789-00",
		Email: "joao@example.com",
	}

	// Configura o mock para retornar o cliente quando GetClientByCPF for chamado
	mockService.On("GetClientByCPF", "123.456.789-00").Return(&client, nil)

	// Cria a requisição GET para buscar o cliente por CPF
	req, _ := http.NewRequest("GET", "/clients/cpf/123.456.789-00", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica se o status da resposta é 200 (OK)
	assert.Equal(t, http.StatusOK, w.Code)

	// Decodifica a resposta JSON e verifica se os dados do cliente estão corretos
	var response clients.Client
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "João Silva", response.Name)
	assert.Equal(t, "123.456.789-00", response.CPF)

	// Verifica se o método GetClientByCPF foi chamado com o CPF correto
	mockService.AssertCalled(t, "GetClientByCPF", "123.456.789-00")
}

// TestGetClientsByName_Success testa a busca de clientes por nome com sucesso.
func TestGetClientsByName_Success(t *testing.T) {
	mockService := new(MockService)
	router := setupRouter(mockService)

	// Define a lista de clientes que será retornada pelo mock
	clientsList := []clients.Client{
		{
			ID:    1,
			Name:  "João Silva",
			Email: "joao@example.com",
		},
		{
			ID:    2,
			Name:  "João Santos",
			Email: "joao.santos@example.com",
		},
	}

	// Configura o mock para retornar a lista de clientes quando GetClientByName for chamado
	mockService.On("GetClientByName", "João").Return(clientsList, nil)

	// Cria a requisição GET para buscar clientes por nome
	req, _ := http.NewRequest("GET", "/clients/name/João", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica se o status da resposta é 200 (OK)
	assert.Equal(t, http.StatusOK, w.Code)

	// Decodifica a resposta JSON e verifica se os dados dos clientes estão corretos
	var response []clients.Client
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 2, len(response))
	assert.Equal(t, "João Silva", response[0].Name)
	assert.Equal(t, "João Santos", response[1].Name)

	// Verifica se o método GetClientByName foi chamado com o nome correto
	mockService.AssertCalled(t, "GetClientByName", "João")
}

// TestGetClientsByName_NotFound testa a busca de clientes por nome inexistente.
func TestGetClientsByName_NotFound(t *testing.T) {
	mockService := new(MockService)
	router := setupRouter(mockService)

	// Configura o mock para retornar uma lista vazia quando GetClientByName for chamado
	mockService.On("GetClientByName", "Inexistente").Return([]clients.Client{}, nil)

	// Cria a requisição GET para buscar clientes por nome
	req, _ := http.NewRequest("GET", "/clients/name/Inexistente", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica se o status da resposta é 200 (OK) e se a lista de clientes está vazia
	assert.Equal(t, http.StatusOK, w.Code)
	var response []clients.Client
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Empty(t, response)

	// Verifica se o método GetClientByName foi chamado com o nome correto
	mockService.AssertCalled(t, "GetClientByName", "Inexistente")
}