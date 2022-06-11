package main

import (
  "github.com/gin-gonic/gin" 
  "github.com/cloudguruab/prospect/tmp"
  "github.com/cloudguruab/prospect/src"
)

// todo:
// add authentication
// group routes
// add cloud function triggers
// setup CI 
// add Testcases
// add logging
// update newfile-service route
// add routes for uploading files
func main() {
    r := gin.Default()
    tmp.ConnectDB()

    r.GET("/api/v1/prospect/", src.Welcome)
    r.POST("/api/v1/prospect/upload", src.Upload)

    r.Run(":8000")
}

