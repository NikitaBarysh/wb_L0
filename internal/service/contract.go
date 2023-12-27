package service

import (
	"github.com/NikitaBarysh/wb_L0/cmd/cfg"
	"github.com/NikitaBarysh/wb_L0/internal/cache"
	"github.com/NikitaBarysh/wb_L0/internal/entity"
	"github.com/NikitaBarysh/wb_L0/internal/repository"
)

type Order interface {
	GetOrder(id string) (entity.Order, error)
	RestoreCache() error
}

type Nats interface {
	SubscribeToNATS() error
}

type NATSProducer interface {
	Publish() error
}

type Service struct {
	Order
	Nats
	NATSProducer
}

func NewService(cache *cache.Cache, cfg *cfg.Config, rep *repository.Repository) *Service {
	return &Service{
		Order:        NewOrderService(cache, rep),
		Nats:         NewNatsService(cache, cfg, rep),
		NATSProducer: NewPublisher(cfg.Channel),
	}
}
