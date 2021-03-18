package main

import (
	"github.com/suadev/microservices/services.product/src/entity"
	product "github.com/suadev/microservices/services.product/src/internal"
	"github.com/suadev/microservices/services.product/src/kafka"
	"github.com/suadev/microservices/shared/config"
	"github.com/suadev/microservices/shared/server"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	config := config.LoadConfig(".")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.GetDBURL(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to the DB.")
	}

	db.AutoMigrate(&entity.Product{})

	repo := product.NewRepository(db)
	service := product.NewService(repo)
	handler := product.NewHandler(service)

	go kafka.RegisterConsumer(config.KafkaOrderTopic, service)

	err = server.NewServer(handler.Init(), config.AppPort).Run()
	if err != nil {
		panic("Couldn't start the HTTP server.")
	}
}
