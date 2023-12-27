package service

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/NikitaBarysh/wb_L0/internal/entity"
	"github.com/bxcodec/faker/v3"
	"github.com/nats-io/stan.go"
)

const (
	cluster = "test-cluster"
	client  = "client_2"
	url     = "nats://localhost:4223"
)

type Producer struct {
	channel string
}

func NewPublisher(channel string) *Producer {
	return &Producer{
		channel: channel,
	}
}

func (s *Producer) Publish() error {
	ns, err := stan.Connect(cluster, client, stan.NatsURL(url))
	if err != nil {
		return fmt.Errorf("err to connect: %w", err)
	}

	for {
		data, err := s.generateRandomOrder()
		if err != nil {
			log.Println("err generateRandomOrder: %w", err)
			continue
		}

		err = ns.Publish(s.channel, data)
		if err != nil {
			log.Println("err to publish: ", err)
			continue
		}
		time.Sleep(time.Second * 5)
	}
}

func (s *Producer) generateRandomOrder() ([]byte, error) {
	var randomOrder entity.Order

	err := faker.FakeData(&randomOrder)
	if err != nil {
		return nil, fmt.Errorf("err to do fake data: %w", err)
	}

	randomOrder.OrderUid = faker.UUIDDigit()
	randomOrder.TrackNumber = faker.UUIDDigit()
	randomOrder.Entry = "WBIL"
	randomOrder.Locale = "en"
	randomOrder.InternalSignature = ""
	randomOrder.CustomerID = faker.Username()
	randomOrder.DeliveryService = "meest"
	randomOrder.Shardkey = fmt.Sprintf("%d", rand.Intn(10))
	randomOrder.SmID = rand.Intn(100)
	randomOrder.DateCreated = time.Now()
	randomOrder.OofShard = fmt.Sprintf("%d", rand.Intn(10))

	randomOrder.Delivery = s.generateRandomDelivery()
	randomOrder.Payment = s.generateRandomPayment()
	randomOrder.Items = []entity.Item{s.generateRandomItem()}

	data, err := json.Marshal(randomOrder)
	if err != nil {
		return nil, fmt.Errorf("err to marshal: %w", err)
	}

	return data, nil
}

func (s *Producer) generateRandomDelivery() entity.Delivery {
	return entity.Delivery{
		Name:    faker.FirstName(),
		Phone:   faker.Phonenumber(),
		Zip:     fmt.Sprintf("%05d", rand.Intn(100000)),
		City:    "Moscow",
		Address: "red square",
		Region:  "Moscow",
		Email:   faker.Email(),
	}
}

func (s *Producer) generateRandomPayment() entity.Payment {
	return entity.Payment{
		Transaction:  faker.UUIDDigit(),
		RequestID:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       rand.Intn(1000),
		PaymentDT:    rand.Intn(1000),
		Bank:         faker.Word(),
		DeliveryCost: rand.Intn(2000),
		GoodsTotal:   rand.Intn(500),
		CustomFee:    rand.Intn(100),
	}
}

func (s *Producer) generateRandomItem() entity.Item {
	return entity.Item{
		ChrtID:      rand.Intn(1000000),
		TrackNumber: faker.UUIDDigit(),
		Price:       rand.Intn(100),
		RID:         faker.UUIDDigit(),
		Name:        faker.Word(),
		Sale:        rand.Intn(50),
		Size:        fmt.Sprintf("%d", rand.Intn(10)),
		TotalPrice:  rand.Intn(100),
		NmID:        rand.Intn(1000000),
		Brand:       faker.Word(),
		Status:      rand.Intn(300),
	}
}
