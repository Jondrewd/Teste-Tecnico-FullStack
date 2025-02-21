package clients

import (
	"strings"

	"gorm.io/gorm"
)

// Repository é uma interface que define os métodos necessários para operações de banco de dados relacionadas a clientes.
// Essa interface permite que diferentes implementações de repositório sejam usadas, facilitando testes e manutenção.
type Repository interface {
	CreateClient(client *Client) (*Client, error)      // Cria um novo cliente
	GetClients() ([]Client, error)                    // Retorna todos os clientes
	GetClientByID(id uint) (*Client, error)           // Retorna um cliente pelo ID
	UpdateClient(id uint, client *Client) (*Client, error) // Atualiza os dados de um cliente
	DeleteClient(id uint) error                       // Deleta um cliente pelo ID
	FindByCPF(cpf string) (*Client, error)            // Busca um cliente pelo CPF
	CountClients() (int64, error)                     // Retorna o total de clientes cadastrados
	FindByName(name string) ([]Client, error)
}

// repository é uma struct que implementa a interface Repository.
// Ela contém uma instância do GORM (*gorm.DB) para interagir com o banco de dados.
type repository struct {
	db *gorm.DB
}

// NewRepository é uma função que cria e retorna uma nova instância do repositório.
// Recebe uma conexão com o banco de dados (*gorm.DB) e retorna um objeto que implementa a interface Repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// CreateClient cria um novo cliente no banco de dados.
// Recebe um ponteiro para um objeto Client e o persiste no banco de dados usando o GORM.
// Retorna o cliente criado ou um erro, caso ocorra algum problema.
func (r *repository) CreateClient(client *Client) (*Client, error) {
	if err := r.db.Create(client).Error; err != nil {
		return nil, err
	}
	return client, nil
}

// GetClients retorna uma lista de todos os clientes cadastrados no banco de dados.
// Usa o método Find do GORM para buscar todos os registros da tabela de clientes.
// Retorna a lista de clientes ou um erro, caso ocorra algum problema.
func (r *repository) GetClients() ([]Client, error) {
	var clients []Client
	if err := r.db.Find(&clients).Error; err != nil {
		return nil, err
	}
	if err := r.db.Preload("Deliveries").Find(&clients).Error; err != nil {
        return nil, err
		}
	return clients, nil
}

// GetClientByID retorna um cliente específico com base no ID fornecido.
// Usa o método First do GORM para buscar o cliente pelo ID.
// Retorna o cliente encontrado ou um erro, caso o cliente não exista ou ocorra algum problema.
func (r *repository) GetClientByID(id uint) (*Client, error) {
	var client Client
	if err := r.db.First(&client, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Preload("Deliveries").Find(&client).Error; err != nil {
        return nil, err
		}
	return &client, nil
}

// UpdateClient atualiza os dados de um cliente existente no banco de dados.
// Primeiro, busca o cliente pelo ID para garantir que ele existe.
// Em seguida, usa o método Updates do GORM para aplicar as alterações.
// Retorna o cliente atualizado ou um erro, caso ocorra algum problema.
func (r *repository) UpdateClient(id uint, client *Client) (*Client, error) {
	var existingClient Client
	if err := r.db.First(&existingClient, id).Error; err != nil {
		return nil, err
	}

	// Atualiza os campos do cliente existente com os dados fornecidos.
	if err := r.db.Model(&existingClient).Updates(client).Error; err != nil {
		return nil, err
	}

	return &existingClient, nil
}

// DeleteClient remove um cliente do banco de dados com base no ID fornecido.
// Usa o método Delete do GORM para excluir o registro.
// Retorna um erro, caso ocorra algum problema durante a exclusão.
func (r *repository) DeleteClient(id uint) error {
	if err := r.db.Delete(&Client{}, id).Error; err != nil {
		return err
	}
	return nil
}

// FindByCPF busca um cliente no banco de dados com base no CPF fornecido.
// Usa o método Where do GORM para filtrar os registros pelo CPF.
// Retorna o cliente encontrado ou um erro, caso o cliente não exista ou ocorra algum problema.
func (r *repository) FindByCPF(cpf string) (*Client, error) {
	var client Client
	if err := r.db.Where("cpf = ?", cpf).First(&client).Error; err != nil {
		return nil, err
	}
	if err := r.db.Preload("Deliveries").Find(&client).Error; err != nil {
        return nil, err
		}
	return &client, nil
}

// FindByName busca múltiplos clientes no banco de dados com base em uma correspondência parcial do nome.
// Usa o método Where do GORM com LIKE para filtrar os registros pelo nome, garantindo que o nome comece com o termo fornecido.
// A busca não será sensível a maiúsculas/minúsculas.
func (r *repository) FindByName(name string) ([]Client, error) {
	var clients []Client

	// Converte o nome para minúsculas e adiciona o operador LIKE
	searchTerm := strings.ToLower(name) + "%"

	// Realiza a busca com a cláusula WHERE e o Preload em uma única consulta
	if err := r.db.
		Where("LOWER(name) LIKE ?", searchTerm). // Filtra pelo nome
		Preload("Deliveries").                  // Carrega as entregas associadas
		Find(&clients).                         // Executa a consulta
		Error; err != nil {
		return nil, err
	}

	return clients, nil
}


// CountClients retorna o número total de clientes cadastrados no banco de dados.
// Usa o método Count do GORM para contar os registros na tabela de clientes.
// Retorna o total de clientes ou um erro, caso ocorra algum problema.
func (r *repository) CountClients() (int64, error) {
	var count int64
	if err := r.db.Model(&Client{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}