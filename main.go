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

	// smtpPort, err := strconv.Atoi(env.GetEnv("SMTP_PORT"))
	// if err != nil {
	// 	log.Fatalf("Error trying to connect to SMTP port, error: %s\n", err.Error())
	// 	return
	// }

	// mailService := mailer.NewSMTPMailService(mailer.SMTPConfig{
	// 	Host:     env.GetEnv("SMTP_HOST"),
	// 	Port:     smtpPort,
	// 	UserName: env.GetEnv("SMTP_USERNAME"),
	// 	Password: env.GetEnv("SMTP_PASSWORD"),
	// 	From:     env.GetEnv("SMTP_FROM"),
	// })

	// mailService.Send(mailer.MailMessage{
	// 	To:      []string{"jhon.doe@example.com"},
	// 	Subject: "Test email",
	// 	Body:    []byte("This is a test message"),
	// })

	database, err := postgres.NewPostgresConnection()
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL, error: %s \n", err.Error())
		return
	}
	defer database.Close()

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":3333"); err != nil {
		log.Fatal(err)
	}
}
