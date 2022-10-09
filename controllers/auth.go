package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/bianucci/hello-go/models"
	"github.com/bianucci/hello-go/utils"
	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

var users = map[string]models.User{}

func asSha(s string) string {
	sha := sha256.New()
	sha.Write([]byte(s))
	return base64.URLEncoding.EncodeToString(sha.Sum(nil))
}

func Register(c *gin.Context) {

	var input Credentials

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, exists := users[input.Username]
	if exists {
		c.String(http.StatusNotFound, fmt.Sprintf("User already registered=\"%s\"", input.Username))
		return
	}

	var newUser models.User
	newUser.Username = input.Username
	newUser.Password = asSha(input.Password)

	users[input.Username] = newUser

	c.String(http.StatusCreated, "OK")
}

func Login(c *gin.Context) {
	var input Credentials

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user, ok := users[input.Username]; ok {
		token, err := utils.GenerateToken(user.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.String(http.StatusOK, token)
		}
		return
	}

	c.String(http.StatusUnauthorized, "Error!")
}
