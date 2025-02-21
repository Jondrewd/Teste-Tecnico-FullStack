package clients

// Service é uma interface que define os métodos necessários para a camada de serviço de clientes.
// Ela atua como um contrato para a lógica de negócio relacionada a clientes.
type Service interface {
	CreateClient(client *Client) (*Client, error)      // Cria um novo cliente
	GetClients() ([]Client, error)                    // Retorna todos os clientes
	GetClientByID(id uint) (*Client, error)           // Retorna um cliente pelo ID
	UpdateClient(id uint, client *Client) (*Client, error) // Atualiza os dados de um cliente
	DeleteClient(id uint) error                       // Deleta um cliente pelo ID
	GetClientByCPF(cpf string) (*Client, error)       // Busca um cliente pelo CPF
	GetTotalClients() (int64, error)                  // Retorna o número total de clientes
	GetClientByName(name string) ([]Client, error)
}

// service é uma struct que implementa a interface Service.
// Ela contém uma instância de um repositório (Repository) para interagir com a camada de dados.
type service struct {
	repo Repository
}

// NewService é uma função que cria e retorna uma nova instância do serviço de clientes.
// Recebe um repositório (Repository) como dependência e retorna um objeto que implementa a interface Service.
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// CreateClient implementa a lógica para criar um novo cliente.
// Ele delega a operação para o repositório (Repository) e retorna o cliente criado ou um erro.
func (s *service) CreateClient(client *Client) (*Client, error) {
	return s.repo.CreateClient(client)
}

// GetClients implementa a lógica para retornar todos os clientes cadastrados.
// Ele delega a operação para o repositório (Repository) e retorna a lista de clientes ou um erro.
func (s *service) GetClients() ([]Client, error) {
	return s.repo.GetClients()
}

// GetClientByID implementa a lógica para buscar um cliente pelo ID.
// Ele delega a operação para o repositório (Repository) e retorna o cliente encontrado ou um erro.
func (s *service) GetClientByID(id uint) (*Client, error) {
	return s.repo.GetClientByID(id)
}

// UpdateClient implementa a lógica para atualizar os dados de um cliente existente.
// Ele delega a operação para o repositório (Repository) e retorna o cliente atualizado ou um erro.
func (s *service) UpdateClient(id uint, client *Client) (*Client, error) {
	return s.repo.UpdateClient(id, client)
}

// DeleteClient implementa a lógica para deletar um cliente pelo ID.
// Ele delega a operação para o repositório (Repository) e retorna um erro, caso ocorra algum problema.
func (s *service) DeleteClient(id uint) error {
	return s.repo.DeleteClient(id)
}

// GetClientByCPF implementa a lógica para buscar um cliente pelo CPF.
// Ele delega a operação para o repositório (Repository) e retorna o cliente encontrado ou um erro.
func (s *service) GetClientByCPF(cpf string) (*Client, error) {
	return s.repo.FindByCPF(cpf)
}
// GetClientByName implementa a lógica para buscar um cliente pelo Nome.
// Ele delega a operação para o repositório (Repository) e retorna o cliente encontrado ou um erro.
func (s *service) GetClientByName(name string) ([]Client, error) {
	return s.repo.FindByName(name)
}

// GetTotalClients implementa a lógica para retornar o número total de clientes cadastrados.
// Ele delega a operação para o repositório (Repository) e retorna o total de clientes ou um erro.
func (s *service) GetTotalClients() (int64, error) {
	return s.repo.CountClients()
}