package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/luisbarufi/my-money-api/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file: ")
	}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":3333"); err != nil {
		log.Fatal(err)
	}
}
