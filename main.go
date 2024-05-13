package main

import (
	"github.com/alerebal/go-rest-api/db"
	"github.com/alerebal/go-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	port := ":8080"
	routes.RegisterRoutes(server)
	server.Run(port)
}
