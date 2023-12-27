package cache

import (
	"fmt"
	"sync"

	"github.com/NikitaBarysh/wb_L0/internal/entity"
)

type Cache struct {
	CacheData map[string]entity.Order
	mu        sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		CacheData: make(map[string]entity.Order),
	}
}

func (s *Cache) UpdateCache(orders []entity.Order) {
	for _, order := range orders {
		s.CacheData[order.OrderUid] = order
	}
}

func (s *Cache) UpdateOrder(order entity.Order) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.CacheData[order.OrderUid] = order
}

func (s *Cache) GetCacheByID(id string) (entity.Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	order, ok := s.CacheData[id]
	if ok {
		return order, nil
	}
	return entity.Order{}, fmt.Errorf("order not found")
}
