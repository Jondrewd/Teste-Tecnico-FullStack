package deliveries

// @description Dados da entrega
// @type object
type Delivery struct {
    ID           uint    `json:"id" gorm:"primaryKey"`
    ClientCPF    string  `json:"client_cpf" gorm:"not null;index"`
    ClientName   string  `json:"client_name" gorm:"not null"`
    TestName     string  `json:"test_name" gorm:"not null"`
    Weight       float64 `json:"weight" gorm:"not null"`
    Logradouro   string  `json:"logradouro" gorm:"not null"`
    Numero       string  `json:"numero" gorm:"not null"`
    Bairro       string  `json:"bairro" gorm:"not null"`
    Complemento  string  `json:"complemento" gorm:"not null"`
    Cidade       string  `json:"cidade" gorm:"not null"`
    Estado       string  `json:"estado" gorm:"not null"`
    Pais         string  `json:"pais" gorm:"not null"`
    Latitude     float64 `json:"latitude" gorm:"not null"`
    Longitude    float64 `json:"longitude" gorm:"not null"`
    OrderStatus  string  `json:"order_status" gorm:"not null"`
}

const (
	OrderStatusPending   = "Pendente"
	OrderStatusShipped   = "Enviado"
	OrderStatusDelivered = "Entregue"
	OrderStatusCanceled  = "Cancelado"
)
