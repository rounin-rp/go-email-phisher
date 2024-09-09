package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all the routes for the handlers
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	userHandler := NewUserHandler(db)
	emailHandler := NewEmailHandler(db)

	// Define all routes here
	router.GET("/user-links", userHandler.GetAllUserLinks)
	router.POST("/user-link", userHandler.SetUserLink)
	router.GET("/user/:id", userHandler.LinkOpened)
	// Add more routes as needed

	router.POST("/emails", emailHandler.CreateEmailTemplate)
	router.GET("/emails", emailHandler.GetAllEmailTemplates)
	router.POST("/send-emails", emailHandler.SendEmailToUsers)
}
