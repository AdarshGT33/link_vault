package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"main.go/handlers"
	"main.go/models"
)

func RegisterUser(c *gin.Context) {
	var user models.Users

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	if user.Username == "" && user.Email == "" && user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "Required fields were empty"})
		return
	}

	userfound := handlers.DB.Where("email = ?", user.Email).Find(&user)

	if userfound.Error != nil && userfound.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, "Error checking for user")
		return
	}

	if userfound.RowsAffected > 0 {
		c.JSON(http.StatusConflict, "Email is already registered")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	createUser := models.Users{
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	handlers.DB.Create(&createUser)

	c.JSON(http.StatusCreated, gin.H{"data": createUser})

}
