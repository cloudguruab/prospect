package src

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

// authentication will go here
// will be responsible for opening newfile-service.go
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"content": "Welcome"})
}

