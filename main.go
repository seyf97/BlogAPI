package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/seyf97/BlogAPI/db"
	"github.com/seyf97/BlogAPI/routes"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		panic("PORT has to exist in the .env file.")
	}

	addr := "127.0.0.1:" + port
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(addr)
}
