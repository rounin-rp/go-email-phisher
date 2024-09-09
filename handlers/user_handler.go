package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rounin-rp/email-phisher/models"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

type CreateLinkRequest struct {
	UserId string `json:"user_id" binding:"required"`
	Email  string `json:"email" binding:"required"`
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

// SetUsers godoc
// @Summary Set User Links
// @Description set user link
// @Tags users
// @Accept  json
// @Produce  json
// @Param link body CreateLinkRequest true "Link data"
// @Success 201 {object} models.Links
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /user-link [post]
func (h *UserHandler) SetUserLink(c *gin.Context) {
	var userRequest CreateLinkRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save to database
	link := models.Links{
		MagicLink: uuid.New().String(),
		HasOpened: false,
		UserId:    userRequest.UserId,
		Email:     userRequest.Email,
	}
	if err := h.db.Create(&link).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, link)
}

// GetUsers godoc
// @Summary List all users links
// @Description Get all users links
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Links
// @Router /user-links [get]
func (h *UserHandler) GetAllUserLinks(c *gin.Context) {
	var userLinks []models.Links
	err := h.db.Find(&userLinks).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userLinks)
}

// LinkOpened godoc
// @Summary Set has_opened to true
// @Description Set has_opened to true
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "Link ID"
// @Success 200  {object} models.Links
// @Router /user/{id} [get]
func (h *UserHandler) LinkOpened(c *gin.Context) {
	id := c.Param("id")

	log.Println("got the magic link : ", id)
	var link models.Links

	err := h.db.Where("magic_link = ? ", id).First(&link).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	link.HasOpened = true
	link.TimesClicked += 1
	h.db.Save(link)

	c.JSON(http.StatusOK, link)
}
