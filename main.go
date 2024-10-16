package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/luisbarufi/my-money-api/src/configuration/env"
)

func main() {
	godotenv.Load()
	fmt.Println(env.GetEnv("API_PORT"))
}
