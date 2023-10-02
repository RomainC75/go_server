package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RomainC75/postgres-test/models"
	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func SignupUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var foundUser models.User
	if err := models.GetUserByUsername(&foundUser, user.Username); err == nil {
		fmt.Print("FOUND USER : ", foundUser)
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is already used !"})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashed)

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

	// isValid := comparePasswords()

	fmt.Printf("\n-> login User : %+v // found user : %+v \n", loginUser, foundUser)
}
