package main

import (
	"log"

	"github.com/alerebal/go-rest-api/db"
	"github.com/alerebal/go-rest-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	db.InitDB()
	server := gin.Default()
	port := ":8080"
	routes.RegisterRoutes(server)
	server.Run(port)
}
