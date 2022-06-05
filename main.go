package main

import (
  "github.com/gin-gonic/gin" 
  "github.com/cloudguruab/prospect/tmp"
  "github.com/cloudguruab/prospect/src"
)


func main() {
    r := gin.Default()
    tmp.ConnectDB()

    // start grouping routes into gateway
    //
    r.GET("/api/v1/prospect/", src.Welcome)
    r.POST("/api/v1/prospect/upload", src.Upload)

    r.Run(":8000")
}

