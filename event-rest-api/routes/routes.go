package routes

import (
	"event-rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	// Apply authentication middleware to all event routes
	protectedRoutes := server.Group("/").Use(middlewares.Authenticate)
	protectedRoutes.POST("/events", createEvent)
	protectedRoutes.PUT("/events/:id", updateEvent)
	protectedRoutes.DELETE("/events/:id", deleteEvent)
	protectedRoutes.POST("/events/:id/register", registerForEvent)
	protectedRoutes.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
