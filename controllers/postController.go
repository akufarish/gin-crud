package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostIndex (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Ok")
}

func PostStore (c *gin.Context) {

}

func PostSingle (c *gin.Context) {

}

func PostUpdate (c *gin.Context) {

}

func PostDelete (c *gin.Context) {
	
}