package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RomainC75/postgres-test/models"
	"github.com/gin-gonic/gin"
)

func SignupUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if _, err := models.CreateUser(&user); err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, user)
}

func LoginUser(c *gin.Context) {
	var loginUser models.LoginUser
	if err := c.BindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unprocessable"})
		return
	}

	var foundUser models.User
	if err := models.GetUserByUsername(&foundUser, loginUser.Username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	fmt.Printf("\n-> login User : %+v // found user : %+v \n", loginUser, foundUser)
}
