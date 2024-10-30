package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/luisbarufi/my-money-api/src/configuration/database/postgres"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/controller"
	"github.com/luisbarufi/my-money-api/src/controller/routes"
	"github.com/luisbarufi/my-money-api/src/model/repository"
	"github.com/luisbarufi/my-money-api/src/model/service"
)

func main() {
	logger.Info("Starting application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file: ")
	}

	database, err := postgres.NewPostgresConnection()
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL, error: %s \n", err.Error())
		return
	}
	defer database.Close()

	logger.Info("PostgreSQL connection established successfully!")
	logger.Info("Migrations successfully implemented!")

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":3333"); err != nil {
		log.Fatal(err)
	}
}
