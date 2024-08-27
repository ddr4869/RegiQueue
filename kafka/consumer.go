package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ddr4869/RegiQueue/redis"
	"github.com/ddr4869/RegiQueue/service"

	"github.com/IBM/sarama"
)

var consumer sarama.Consumer

func InitConsumer() error {
	var err error
	consumer, err = sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatal("Error creating consumer: ", err)
	}
	return nil
}

func ConsumeMessages(topic string) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("Error creating partition consumer: ", err)
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		var req service.RegistrationRequest
		err := json.Unmarshal(message.Value, &req)
		if err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		// Process registration
		success := service.ProcessRegistration(context.Background(), req)

		if !success {
			log.Printf("Registration failed for user: %s, course: %s", req.UserID, req.CourseName)
		}

		// Remove user from queue
		redis.DecrementQueue(context.Background(), req.UserID)
	}
}

func CloseConsumer() {
	consumer.Close()
}
