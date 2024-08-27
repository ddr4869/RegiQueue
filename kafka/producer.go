package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

var producer sarama.SyncProducer

func InitProducer() error {
	var err error
	producer, err = sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatal("Error creating consumer: ", err)
	}
	return nil
}

// func SendMessage(topic, message string) error {
// 	msg := &sarama.ProducerMessage{
// 		Topic: topic,
// 		Value: sarama.StringEncoder(message),
// 		//Value: sarama.ByteEncoder(message),
// 	}
// 	var err error
// 	_, _, err = producer.SendMessage(msg)
// 	if err != nil {
// 		log.Printf("Failed to send message: %v", err)
// 		return err
// 	}

// 	log.Printf("Message sent to topic %s: %s", topic, message)
// 	return nil
// }

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
