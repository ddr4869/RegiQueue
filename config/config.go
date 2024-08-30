package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB    DbConfig
	Gin   GinConfig
	Kafka KafkaConfig
	Redis RedisConfig
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

type GinConfig struct {
	Mode string
}

type KafkaConfig struct {
	Topic           string
	ProducerAddress string
	ConsumerAddress string
}

type RedisConfig struct {
	Address  string
	Password string
}

func Init() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GinConfig := GinConfig{
		Mode: os.Getenv("GIN_MODE"),
	}

	DbConfig := DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Name:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	KafkaConfig := KafkaConfig{
		Topic:           os.Getenv("KAFKA_TOPIC"),
		ProducerAddress: os.Getenv("KAFKA_PRODUCER_ADDRESS"),
		ConsumerAddress: os.Getenv("KAFKA_CONSUMER_ADDRESS"),
	}

	RedisConfig := RedisConfig{
		Address:  os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	return &Config{
		Gin:   GinConfig,
		DB:    DbConfig,
		Kafka: KafkaConfig,
		Redis: RedisConfig,
	}
}
