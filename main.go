package main

import (
  "github.com/gin-gonic/gin" 
  // "github.com/cloudguruab/prospect/tmp"
  "github.com/cloudguruab/prospect/src"
)


func main() {
    r := gin.Default()

    // tmp.ConnectDB()


    r.GET("/", src.Welcome)

    r.Run()
}

