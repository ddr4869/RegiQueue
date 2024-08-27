package main

import (
	"log"
	"net/http"

	"github.com/ddr4869/RegiQueue/handlers"
	"github.com/ddr4869/RegiQueue/kafka"
	"github.com/ddr4869/RegiQueue/redis"
)

func main() {
	// Redis 초기화
	redis.InitRedis()

	// Kafka 컨슈머를 별도의 고루틴으로 시작
	kafka.InitProducer()
	kafka.InitConsumer()
	defer kafkaClose()
	go kafka.ConsumeMessages("registration_topic")

	// HTTP 핸들러 설정
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/queue_position", handlers.GetQueuePosition)
	http.HandleFunc("/run_load_test", handlers.RunLoadTest)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func kafkaClose() {
	kafka.CloseConsumer()
	kafka.CloseProducer()
}
