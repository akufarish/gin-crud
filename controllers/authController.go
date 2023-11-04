package controllers

import (
	"net/http"

	"github.com/akufarish/gin-crud/db"
	"github.com/akufarish/gin-crud/forms"
	"github.com/akufarish/gin-crud/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.Bind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}


	if db.DB.Where("email = ?", user.Email).First(&user) == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Account already registered",
		})
		c.Abort()
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	user.Password = string(hash)

	db.DB.Create(&user)

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Register successfull",
		"user": user,
	})
}

func Login(c *gin.Context)  {
	var user models.User
	var payload forms.LoginPayload

	if err := c.Bind(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	db.DB.Where("email = ?", payload.Email).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "incorrect password",
		})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Login successfull",
		"user": user,
	})

}