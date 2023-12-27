package service

import (
	"fmt"

	"github.com/NikitaBarysh/wb_L0/internal/cache"
	"github.com/NikitaBarysh/wb_L0/internal/entity"
	"github.com/NikitaBarysh/wb_L0/internal/repository"
)

type OrderService struct {
	rep   repository.Order
	cache *cache.Cache
}

func NewOrderService(cache *cache.Cache, newRep *repository.Repository) *OrderService {
	return &OrderService{
		rep:   newRep,
		cache: cache,
	}
}

func (s *OrderService) RestoreCache() error {
	orders, err := s.rep.GetAllOrders()
	if err != nil {
		return fmt.Errorf("err to get orders: %w", err)
	}
	s.cache.UpdateCache(orders)
	return nil
}

func (s *OrderService) GetOrder(id string) (entity.Order, error) {
	order, err := s.cache.GetCacheByID(id)
	if err != nil {
		return entity.Order{}, fmt.Errorf("err to GetOrder: %w", err)
	}

	return order, nil
}
