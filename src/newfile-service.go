package src

import (
  "github.com/gin-gonic/gin" 
  "github.com/cloudguruab/prospect/tmp"
    
  "net/http"
)


func NewFileService() {
    r := gin.Default()

    tmp.ConnectDB()


    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
    })

    r.Run()
}

