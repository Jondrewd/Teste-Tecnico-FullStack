package deliveries

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

// Repository é uma interface que define os métodos que o repositório deve implementar.
// Ela serve como um contrato para a camada de acesso a dados relacionada a entregas.
type Repository interface {
	CreateDelivery(delivery *Delivery) (*Delivery, error) // Cria uma nova entrega
	GetDeliveries() ([]Delivery, error)                  // Retorna todas as entregas
	GetDeliveryByID(id uint) (*Delivery, error)          // Retorna uma entrega pelo ID
	UpdateDelivery(id uint, delivery *Delivery) (*Delivery, error) // Atualiza uma entrega
	DeleteDelivery(id uint) error                        // Deleta uma entrega pelo ID
	FindByCPF(cpf string) ([]Delivery, error)            // Busca entregas pelo CPF do cliente
	FindByClientName(name string) ([]Delivery, error) // Busca entregas pelo Nome do cliente
	FindByCity(city string) ([]Delivery, error) // Busca entregas pelo Nome da cidade
	UpdateOrderStatus(id uint, status string) error      // Atualiza o status de uma entrega
}

// repository é uma struct que implementa a interface Repository.
// Ela contém uma instância do GORM (*gorm.DB) para interagir com o banco de dados.
type repository struct {
	db *gorm.DB
}

// NewRepository cria uma nova instância do repositório.
// Recebe uma conexão com o banco de dados (*gorm.DB) e retorna um objeto que implementa a interface Repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// CreateDelivery cria uma nova entrega no banco de dados.
// Recebe um ponteiro para um objeto Delivery e o persiste no banco de dados usando o GORM.
// Retorna a entrega criada ou um erro, caso ocorra algum problema.
func (r *repository) CreateDelivery(delivery *Delivery) (*Delivery, error) {
	if err := r.db.Create(&delivery).Error; err != nil {
		return nil, err
	}
	return delivery, nil
}

// GetDeliveries retorna todas as entregas cadastradas no banco de dados.
// Usa o método Find do GORM para buscar todos os registros da tabela de entregas.
// Retorna a lista de entregas ou um erro, caso ocorra algum problema.
func (r *repository) GetDeliveries() ([]Delivery, error) {
	var deliveries []Delivery
	if err := r.db.Find(&deliveries).Error; err != nil {
		return nil, err
	}
	return deliveries, nil
}

// GetDeliveryByID retorna uma entrega específica com base no ID fornecido.
// Usa o método First do GORM para buscar a entrega pelo ID.
// Retorna a entrega encontrada ou um erro, caso a entrega não exista ou ocorra algum problema.
func (r *repository) GetDeliveryByID(id uint) (*Delivery, error) {
	var delivery Delivery
	if err := r.db.First(&delivery, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("delivery not found")
		}
		return nil, err
	}
	return &delivery, nil
}

// UpdateDelivery atualiza os dados de uma entrega existente no banco de dados.
// Primeiro, busca a entrega pelo ID para garantir que ela existe.
// Em seguida, usa o método Updates do GORM para aplicar as alterações.
// Retorna a entrega atualizada ou um erro, caso ocorra algum problema.
func (r *repository) UpdateDelivery(id uint, delivery *Delivery) (*Delivery, error) {
	var existingDelivery Delivery
	if err := r.db.First(&existingDelivery, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("delivery not found")
		}
		return nil, err
	}

	// Atualiza os campos da entrega existente com os dados fornecidos.
	if err := r.db.Model(&existingDelivery).Updates(delivery).Error; err != nil {
		return nil, err
	}
	return &existingDelivery, nil
}

// DeleteDelivery remove uma entrega do banco de dados com base no ID fornecido.
// Primeiro, verifica se a entrega existe.
// Em seguida, usa o método Delete do GORM para excluir o registro.
// Retorna um erro, caso ocorra algum problema durante a exclusão.
func (r *repository) DeleteDelivery(id uint) error {
	var delivery Delivery
	if err := r.db.First(&delivery, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("delivery not found")
		}
		return err
	}

	if err := r.db.Delete(&delivery).Error; err != nil {
		return err
	}
	return nil
}

// FindByCPF busca todas as entregas associadas a um CPF específico.
// Usa o método Where do GORM para filtrar os registros pelo CPF.
// Retorna a lista de entregas ou um erro, caso ocorra algum problema.
func (r *repository) FindByCPF(cpf string) ([]Delivery, error) {
	var deliveries []Delivery
	if err := r.db.Where("client_cpf = ?", cpf).Find(&deliveries).Error; err != nil {
		return nil, err
	}
	return deliveries, nil
}

// FindByClientName busca múltiplas entregas associadas a um cliente no banco de dados com base em uma correspondência parcial do nome do cliente.
// Usa o método Where do GORM com LIKE para filtrar os registros pelo nome do cliente.
// A busca não será sensível a maiúsculas/minúsculas.
func (r *repository) FindByClientName(name string) ([]Delivery, error) {
	var deliveries []Delivery

	// Converte o nome para minúsculas e adiciona o operador LIKE
	searchTerm := strings.ToLower(name) + "%"

	// Realiza a busca com a cláusula WHERE e o Preload em uma única consulta
	if err := r.db.
		// Realiza a busca no nome do cliente associado à entrega
		Joins("JOIN clients c ON c.name = deliveries.client_name").
		Where("LOWER(c.name) LIKE ?", searchTerm). // Filtra pelo nome do cliente
		Find(&deliveries).                          // Executa a consulta
		Error; err != nil {
		return nil, err
	}

	return deliveries, nil
}

// FindByCity busca múltiplas entregas associadas a uma cidade no banco de dados com base em uma correspondência parcial do nome da cidade.
// Usa o método Where do GORM com LIKE para filtrar os registros pelo nome da cidade.
// A busca não será sensível a maiúsculas/minúsculas.
func (r *repository) FindByCity(city string) ([]Delivery, error) {
	var deliveries []Delivery

	// Converte o nome da cidade para minúsculas e adiciona o operador LIKE
	searchTerm := strings.ToLower(city) + "%"

	// Realiza a busca com a cláusula WHERE
	if err := r.db.
		Where("LOWER(cidade) LIKE ?", searchTerm). // Filtra pelo nome da cidade
		Find(&deliveries).                         // Executa a consulta
		Error; err != nil {
		return nil, err
	}

	return deliveries, nil
}


// UpdateOrderStatus atualiza o status de uma entrega no banco de dados.
// Usa o método Update do GORM para alterar o campo "order_status" da entrega com o ID fornecido.
// Retorna um erro, caso ocorra algum problema durante a atualização.
func (r *repository) UpdateOrderStatus(id uint, status string) error {
	return r.db.Model(&Delivery{}).Where("id = ?", id).Update("order_status", status).Error
}