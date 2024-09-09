package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rounin-rp/email-phisher/models"
	"github.com/rounin-rp/email-phisher/services"
	"gorm.io/gorm"
)

type EmailHandler struct {
	db *gorm.DB
}

type CreateEmailTemplate struct {
	Subject string `json:"subject" binding:"required"`
	Message string `json:"message" binding:"required"`
	Link    string `json:"link"`
}

type UserEmailMap struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
}

type SendEmailToUsersRequest struct {
	TemplateId int            `json:"template_id" binding:"required"`
	UserEmails []UserEmailMap `json:"user_emails" binding:"required"`
}

func NewEmailHandler(db *gorm.DB) *EmailHandler {
	return &EmailHandler{db: db}
}

// SendEmailToUsers godoc
// @Summary Sends the phising mail to a group of users
// @Description sends the phising email provided by template id to a group of users
// @Tags emails
// @Accept  json
// @Produce  json
// @Param link body SendEmailToUsersRequest true "Email data"
// @Success 200 {object} gin.H
// @Failure 402 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /send-emails [post]
func (h *EmailHandler) SendEmailToUsers(c *gin.Context) {
	var request SendEmailToUsersRequest
	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
	sender := os.Getenv("EMAIL_SENDER")
	password := os.Getenv("EMAIL_PASSWORD")
	emailManager := services.BuildEmailManager(host, port, sender, password)

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var emailTemplate models.Email

	err = h.db.Where("id = ?", request.TemplateId).First(&emailTemplate).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	var userLinks []models.Links

	for _, userEmail := range request.UserEmails {
		userLink := models.Links{
			UserId:    userEmail.UserId,
			MagicLink: uuid.New().String(),
			HasOpened: false,
			Email:     userEmail.Email,
		}

		userLinks = append(userLinks, userLink)
	}

	total := len(userLinks)
	success := 0
	// send email and save to database
	for _, link := range userLinks {
		err = h.db.Create(&link).Error

		if err != nil {
			log.Println("Error while saving user link : ", err.Error())
			continue
		}
		completeMessage := emailTemplate.Message + "\n" + emailTemplate.Link + link.MagicLink
		_, err = emailManager.SendMail(link.Email, emailTemplate.Subject, completeMessage)

		if err != nil {
			log.Println("Error while sending email : ", err.Error())
		}
		success += 1
	}

	percentage_success := success * 100 / total
	message := fmt.Sprintf("success rate = %v percent", percentage_success)
	c.JSON(http.StatusAccepted, gin.H{"message": message})
}

// SetEmailTemplate godoc
// @Summary Creates a new email template
// @Description creates a new email template
// @Tags emails
// @Accept  json
// @Produce  json
// @Param link body CreateEmailTemplate true "Email data"
// @Success 201 {object} models.Email
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /emails [post]
func (h *EmailHandler) CreateEmailTemplate(c *gin.Context) {
	var request CreateEmailTemplate

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// save to database

	email := models.Email{
		Subject: request.Subject,
		Message: request.Message,
		Link:    request.Link,
	}

	if err := h.db.Create(&email).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, email)
}

// GetEmaiTemplates godoc
// @Summary List all email templates
// @Description Get all email templates
// @Tags emails
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Email
// @Router /emails [get]
func (h *EmailHandler) GetAllEmailTemplates(c *gin.Context) {
	var emails []models.Email

	err := h.db.Find(&emails).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, emails)
}
