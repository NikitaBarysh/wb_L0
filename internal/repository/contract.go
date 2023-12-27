package repository

import (
	"github.com/NikitaBarysh/wb_L0/internal/entity"
	"gorm.io/gorm"
)

type Order interface {
	SetOrder(order *entity.Order) error
	GetAllOrders() ([]entity.Order, error)
	OrderAutoMigrate() error
}

type Repository struct {
	Order
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Order: NewOrderRepository(db),
	}
}
