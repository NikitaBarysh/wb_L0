package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/NikitaBarysh/wb_L0/cmd/cfg"
	"github.com/NikitaBarysh/wb_L0/internal/cache"
	"github.com/NikitaBarysh/wb_L0/internal/entity"
	"github.com/NikitaBarysh/wb_L0/internal/repository"
	"github.com/nats-io/stan.go"
)

type NatsService struct {
	rep   repository.Order
	cfg   *cfg.Config
	cache *cache.Cache
}

func NewNatsService(cache *cache.Cache, newCfg *cfg.Config, rep *repository.Repository) *NatsService {
	return &NatsService{
		rep:   rep,
		cfg:   newCfg,
		cache: cache,
	}
}

func (s *NatsService) SubscribeToNATS() error {
	sc, err := stan.Connect(s.cfg.ClusterID, s.cfg.ClientID, stan.NatsURL(s.cfg.NATSEndpoint))
	if err != nil {
		return fmt.Errorf("err to connect to nats-streaming: %w", err)
	}
	_, err = sc.Subscribe(s.cfg.Channel, func(msg *stan.Msg) {
		err := s.parseNATSMessage(msg.Data)
		if err != nil {
			log.Println("Error parsing NATS message:", err)
		}
	}, stan.DurableName("test-durable-name"))
	if err != nil {
		return fmt.Errorf("can't subscribe to channel: %w", err)
	}

	return nil
}

func (s *NatsService) parseNATSMessage(data []byte) error {
	var order entity.Order
	err := json.Unmarshal(data, &order)
	if err != nil {
		return fmt.Errorf("err to unmarshal data: %w", err)
	}

	err = s.rep.SetOrder(&order)

	s.cache.UpdateOrder(order)

	return nil
}
