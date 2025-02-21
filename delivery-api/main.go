package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	"delivery-api/config"
	"delivery-api/internal/clients"
	"delivery-api/internal/deliveries"
	_ "delivery-api/docs" // Importa a documentação gerada pelo Swagger
)

// @description Esta é a documentação da API de entregas e clientes utilizando Gin e Swagger.
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
// @contact.name Nome do Contato
// @contact.email example@example.com
// @contact.url http://www.example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

func main() {
	// Inicializa a conexão com o banco de dados.
	db := config.InitDB()

	// Migra as tabelas no banco de dados.
	// Isso garante que as tabelas necessárias para Client e Delivery estejam criadas.
	if err := db.AutoMigrate(&clients.Client{}, &deliveries.Delivery{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Cria as instâncias do repositório e serviço para clientes.
	clientRepo := clients.NewRepository(db)
	clientService := clients.NewService(clientRepo)

	// Cria as instâncias do repositório e serviço para entregas.
	deliveryRepo := deliveries.NewRepository(db)
	deliveryService := deliveries.NewService(deliveryRepo)

	// Cria os handlers para clientes e entregas.
	// Os handlers são responsáveis por lidar com as requisições HTTP.
	clientHandler := clients.Handler{Service: clientService}
	deliveryHandler := deliveries.Handler{Service: deliveryService}

	// Cria uma instância do servidor Gin.
	r := gin.Default()

	// Configura o middleware CORS para permitir requisições de diferentes origens.
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permite todas as origens (altere para segurança em produção)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Métodos HTTP permitidos
		AllowHeaders:     []string{"Content-Type", "Authorization"}, // Cabeçalhos permitidos
		ExposeHeaders:    []string{"Content-Length"}, // Cabeçalhos expostos
		AllowCredentials: true, // Permite credenciais (cookies, autenticação)
		MaxAge:           12 * time.Hour, // Tempo de cache para as configurações do CORS
	}))

	// Rotas para clientes:
	r.POST("/api/v1/clients", clientHandler.CreateClient)          // Cria um novo cliente
	r.GET("/api/v1/clients", clientHandler.GetClients)            // Retorna todos os clientes
	r.GET("/api/v1/clients/cpf/:cpf", clientHandler.GetClientByCPF) // Busca um cliente pelo CPF
	r.GET("/api/v1/clients/:id", clientHandler.GetClientByID)     // Retorna um cliente pelo ID
	r.GET("/api/v1/clients/name/:name", clientHandler.GetClientsByName)     // Retorna um cliente pelo Nome
	r.GET("/api/v1/clients/count", clientHandler.GetTotalClients) // Retorna o total de clientes
	r.PUT("/api/v1/clients/:id", clientHandler.UpdateClient)      // Atualiza um cliente pelo ID
	r.DELETE("/api/v1/clients/:id", clientHandler.DeleteClient)   // Deleta um cliente pelo ID

	// Rotas para entregas:
	r.POST("/api/v1/deliveries", deliveryHandler.CreateDelivery)          // Cria uma nova entrega
	r.GET("/api/v1/deliveries", deliveryHandler.GetDeliveries)           // Retorna todas as entregas
	r.GET("/api/v1/deliveries/:id", deliveryHandler.GetDeliveryByID)     // Retorna uma entrega pelo ID
	r.GET("/api/v1/deliveries/client/cpf/:cpf", deliveryHandler.GetDeliveriesByCPF) // Busca entregas pelo CPF do cliente
	r.GET("/api/v1/deliveries/client/name/:name", deliveryHandler.GetDeliveriesByClientName) // Busca entregas pelo Nome do cliente
	r.GET("/api/v1/deliveries/city/:city", deliveryHandler.GetDeliveriesByCity) // Busca entregas pelo Nome do cliente
	r.PUT("/api/v1/deliveries/:id", deliveryHandler.UpdateDelivery)      // Atualiza uma entrega pelo ID
	r.DELETE("/api/v1/deliveries/:id", deliveryHandler.DeleteDelivery)   // Deleta uma entrega pelo ID
	r.PATCH("/api/v1/deliveries/:id/:status", deliveryHandler.UpdateOrderStatus) // Atualiza o status de uma entrega

	// Rota para o Swagger UI.
	// Acesse http://localhost:8080/swagger/index.html para visualizar a documentação da API.
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Inicia o servidor na porta 8080.
	log.Println("Servidor rodando na porta 8080")
	r.Run(":8080")
}