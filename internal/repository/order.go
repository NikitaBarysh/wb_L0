package repository

import (
	"fmt"

	"github.com/NikitaBarysh/wb_L0/internal/entity"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(newDB *gorm.DB) *OrderRepository {
	rep := &OrderRepository{
		db: newDB,
	}
	return rep
}

func (r *OrderRepository) OrderAutoMigrate() error {
	err := r.db.AutoMigrate(&entity.Order{})
	if err != nil {
		return fmt.Errorf("error AutoMigrate: %v", err)
	}
	return nil
}

func (r *OrderRepository) SetOrder(order *entity.Order) error {

	err := r.db.Create(order).Error
	if err != nil {
		return fmt.Errorf("err to set order: %w", err)
	}
	return nil
}

func (r *OrderRepository) GetAllOrders() ([]entity.Order, error) {
	var orderSlice []entity.Order
	err := r.db.Find(&orderSlice).Error
	if err != nil {
		return nil, fmt.Errorf("err to get orders: %w", err)
	}

	return orderSlice, nil
}
