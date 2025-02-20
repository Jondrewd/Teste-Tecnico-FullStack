package clients

import (
	"delivery-api/internal/deliveries"
	"time"

	"github.com/go-playground/validator/v10"
)

// @description Dados da entrega
// @type object
type Client struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name       string                `json:"name" validate:"required"`
	CPF        string                `json:"cpf" gorm:"unique;not null" validate:"required"`
	CNPJ       string                 `json:"cnpj" gorm:"unique;not null" validate:"required,cnpj"`
	BirthDate  time.Time              `json:"birth_date" validate:"required,birthdate"`
	Email      string                `json:"email" validate:"required,email"`
	Phone      string                `json:"phone" validate:"required"`
	Deliveries []deliveries.Delivery `json:"deliveries" gorm:"foreignKey:ClientCPF;references:CPF"`
}

func (c *Client) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
