package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

var producer sarama.SyncProducer

func InitProducer(address string) error {
	var err error
	producer, err = sarama.NewSyncProducer([]string{address}, nil)
	if err != nil {
		log.Fatal("Error creating consumer: ", err)
	}
	return nil
}

func SendMessage(topic string, data []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		//Value: sarama.StringEncoder(data),
		Value: sarama.ByteEncoder(data),
	}
	var err error
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}

	return nil
}

func CloseProducer() {
	producer.Close()
}
