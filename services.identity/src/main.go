package main

import (
	"github.com/suadev/microservices/services.identity/src/entity"
	identity "github.com/suadev/microservices/services.identity/src/internal"
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

	db.AutoMigrate(&entity.User{})

	repo := identity.NewRepository(db)
	service := identity.NewService(repo)
	handler := identity.NewHandler(service)

	err = server.NewServer(handler.Init(), config.AppPort).Run()
	if err != nil {
		panic("Couldn't start the HTTP server.")
	}
}
