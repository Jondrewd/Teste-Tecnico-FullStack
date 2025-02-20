package deliveries

import (
	"fmt"
)

// Service é uma interface que define os métodos do serviço relacionado a entregas.
// Ela serve como um contrato para a camada de lógica de negócio.
type Service interface {
	CreateDelivery(delivery *Delivery) (*Delivery, error)      // Cria uma nova entrega
	GetDeliveries() ([]Delivery, error)                      // Retorna todas as entregas
	GetDeliveryByID(id uint) (*Delivery, error)              // Retorna uma entrega pelo ID
	UpdateDelivery(id uint, delivery *Delivery) (*Delivery, error) // Atualiza uma entrega
	DeleteDelivery(id uint) error                            // Deleta uma entrega pelo ID
	GetDeliveriesByCPF(cpf string) ([]Delivery, error)       // Busca entregas por CPF
	UpdateOrderStatus(id uint, status string) error          // Atualiza o status de uma entrega
}

// service é uma struct que implementa a interface Service.
// Ela contém uma instância de um repositório (Repository) para interagir com a camada de dados.
type service struct {
	repo Repository
}

// NewService cria uma nova instância de service.
// Recebe um repositório (Repository) como dependência e retorna um objeto que implementa a interface Service.
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// CreateDelivery implementa a lógica para criar uma nova entrega.
// Ele valida o status da entrega antes de delegar a operação para o repositório.
func (s *service) CreateDelivery(delivery *Delivery) (*Delivery, error) {
	// Verifica se o status da entrega é válido.
	if !isValidOrderStatus(delivery.OrderStatus) {
		return nil, fmt.Errorf("invalid order status")
	}

	// Delega a criação da entrega para o repositório.
	return s.repo.CreateDelivery(delivery)
}

// GetDeliveries implementa a lógica para retornar todas as entregas cadastradas.
// Ele delega a operação para o repositório.
func (s *service) GetDeliveries() ([]Delivery, error) {
	return s.repo.GetDeliveries()
}

// GetDeliveryByID implementa a lógica para buscar uma entrega pelo ID.
// Ele delega a operação para o repositório.
func (s *service) GetDeliveryByID(id uint) (*Delivery, error) {
	return s.repo.GetDeliveryByID(id)
}

// UpdateDelivery implementa a lógica para atualizar os dados de uma entrega existente.
// Ele valida o status da entrega antes de delegar a operação para o repositório.
func (s *service) UpdateDelivery(id uint, delivery *Delivery) (*Delivery, error) {
	// Verifica se o status da entrega é válido.
	if !isValidOrderStatus(delivery.OrderStatus) {
		return nil, fmt.Errorf("invalid order status")
	}

	// Delega a atualização da entrega para o repositório.
	return s.repo.UpdateDelivery(id, delivery)
}

// DeleteDelivery implementa a lógica para deletar uma entrega pelo ID.
// Ele delega a operação para o repositório.
func (s *service) DeleteDelivery(id uint) error {
	return s.repo.DeleteDelivery(id)
}

// GetDeliveriesByCPF implementa a lógica para buscar entregas associadas a um CPF específico.
// Ele delega a operação para o repositório.
func (s *service) GetDeliveriesByCPF(cpf string) ([]Delivery, error) {
	return s.repo.FindByCPF(cpf)
}

// UpdateOrderStatus implementa a lógica para atualizar o status de uma entrega.
// Ele valida o novo status antes de delegar a operação para o repositório.
func (s *service) UpdateOrderStatus(id uint, status string) error {
	// Verifica se o novo status é válido.
	if !isValidOrderStatus(status) {
		return fmt.Errorf("invalid order status")
	}

	// Delega a atualização do status para o repositório.
	return s.repo.UpdateOrderStatus(id, status)
}

// isValidOrderStatus verifica se o status da entrega é válido.
// Ele compara o status fornecido com uma lista de status válidos.
func isValidOrderStatus(status string) bool {
	// Lista de status válidos para uma entrega.
	validStatuses := []string{OrderStatusPending, OrderStatusShipped, OrderStatusDelivered, OrderStatusCanceled}

	// Verifica se o status fornecido está na lista de status válidos.
	for _, s := range validStatuses {
		if s == status {
			return true
		}
	}
	return false
}