package src

import (
  "net/http"

  "github.com/gin-gonic/gin"
)


func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Hello world"})
}

