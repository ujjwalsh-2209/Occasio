package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ujjwalsh-2209/event-booking/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    //GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent) // /events1, /events5

	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
