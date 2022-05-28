package main

import (
  "github.com/gin-gonic/gin" 
  "github.com/cloudguruab/prospect/models"
    
  "net/http"
)


func main() {
    r := gin.Default()

    models.ConnectDB()


    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
    })

    r.Run()
}
