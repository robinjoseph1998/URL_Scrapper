package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UrlScrapper(c *gin.Context) {
	var request []string
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
