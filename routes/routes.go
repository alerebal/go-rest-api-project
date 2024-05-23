package routes

import (
	"github.com/alerebal/go-rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticeted := server.Group("/")
	authenticeted.Use(middlewares.Authenticate)
	authenticeted.POST("/events", createEvent)
	authenticeted.PUT("/events/:id", updateEvent)
	authenticeted.DELETE("/events/:id", deleteEvent)
	authenticeted.POST("/events/:id/register", registerForEvent)
	authenticeted.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)

	server.GET("/users", getUsers)
	server.GET("/users/:id", getUser)
	server.DELETE("/users/:id", deleteUser)
}
