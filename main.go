package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/luisbarufi/my-money-api/src/configuration/database/postgres"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	accountController "github.com/luisbarufi/my-money-api/src/controller/accounts"
	accountRoutes "github.com/luisbarufi/my-money-api/src/controller/accounts/routes"
	userController "github.com/luisbarufi/my-money-api/src/controller/users"
	userRoutes "github.com/luisbarufi/my-money-api/src/controller/users/routes"
	accountRepository "github.com/luisbarufi/my-money-api/src/model/accounts/repository"
	accountsService "github.com/luisbarufi/my-money-api/src/model/accounts/service"
	userRepository "github.com/luisbarufi/my-money-api/src/model/users/repository"
	userService "github.com/luisbarufi/my-money-api/src/model/users/service"
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

	userRepo := userRepository.NewUserRepository(database)
	userService := userService.NewUserDomainService(userRepo)
	user := userController.NewUserControllerInterface(userService)

	accountRepo := accountRepository.NewAccountRepository(database)
	accountService := accountsService.NewAccountDomainService(accountRepo)
	account := accountController.NewAccountControllerInterface(accountService)

	router := gin.Default()
	userRoutes.InitRoutes(&router.RouterGroup, user)
	accountRoutes.InitRoutes(&router.RouterGroup, account)

	if err := router.Run(":3333"); err != nil {
		log.Fatal(err)
	}
}
