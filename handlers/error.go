package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Abort(c *gin.Context, err error) {
	log.Println(err)
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
}
