package main

import (
	"log"

	"github.com/ddr4869/RegiQueue/config"
	"github.com/ddr4869/RegiQueue/internal"
	"github.com/ddr4869/RegiQueue/kafka"
	"github.com/ddr4869/RegiQueue/redis"
)

func main() {

	cfg := config.Init()
	redis.InitRedis(cfg.Redis.Address, cfg.Redis.Password)
	kafka.InitProducer(cfg.Kafka.ProducerAddress)
	kafka.InitConsumer(cfg.Kafka.ConsumerAddress)
	defer func() {
		kafka.CloseConsumer()
		kafka.CloseProducer()
	}()

	go kafka.ConsumeMessages(cfg.Kafka.Topic)
	router, err := internal.NewRestController(cfg)
	if err != nil {
		log.Fatalf("failed creating server: %v", err)
	}
	router.Start()
}
