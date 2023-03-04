package main

import (
	"efishery-auth/database"
	"efishery-auth/handler"
	"efishery-auth/router"
	"efishery-auth/service"
	"os"

	"github.com/go-rel/rel"
	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()

	adapter := database.InitDB()
	repo := rel.New(adapter)

	authService := service.NewAuthService(repo, os.Getenv("JWT_SECRET"))

	handler := router.BuildHandler(
		handler.NewAuthHandler(authService),
	)
	router.RunServer(handler)
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
