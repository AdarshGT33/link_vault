package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"main.go/handlers"
	"main.go/models"
)

func LoginUser(c *gin.Context) {
	var user models.Users

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	if user.Email == "" && user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "Fill all the requried fields"})
		return
	}

	var userfound models.Users
	handlers.DB.Where("email = ?", &user.Email).Find(&userfound)

	if userfound.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "User Not Found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userfound.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "Enter correct password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     userfound.ID,
		"expiry": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  userfound.Username,
	})
}
