package deliveries

// @description Dados da entrega
// @type object
type Delivery struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ClientCPF   string  `json:"cpf" gorm:"not null;index"` // Criando um índice para otimizar consultas
	ClientName  string  `json:"client_name" gorm:"not null"`
	TestName    string  `json:"test_name" gorm:"not null"`
	Weight      float64 `json:"weight" gorm:"not null"`
	Address     string  `json:"address" gorm:"not null"`
	Latitude    float64 `json:"latitude" gorm:"not null"`
	Longitude   float64 `json:"longitude" gorm:"not null"`
	OrderStatus string  `json:"order_status" gorm:"not null"`
}

const (
	OrderStatusPending   = "Pending"
	OrderStatusShipped   = "Shipped"
	OrderStatusDelivered = "Delivered"
	OrderStatusCanceled  = "Canceled"
)
