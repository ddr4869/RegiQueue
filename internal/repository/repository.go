package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/ddr4869/RegiQueue/ent"
)

type Repository struct {
	entClient *ent.Client
}

func (r *Repository) NewEntClient(DBHost, DBPort, DBUser, DBName, DBPassword string) error {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DBHost, DBPort, DBUser, DBName, DBPassword)
	client, err := ent.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		return err
	}
	r.entClient = client
	return nil
}
