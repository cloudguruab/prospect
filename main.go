package main

import (
	"github.com/cloudguruab/prospect/src"
	"github.com/gin-gonic/gin"
)

// todo:
// add authentication
// add cloud function triggers
// add data layer
// patch readme for docs
// setup CI
// add Testcases
// add logging
// update newfile-service route
// add routes for uploading files
func main() {
	r := gin.Default()

	// database
	src.ConnectDB()

	// routes
	r.GET("/api/v1/prospect/", src.Index)
	r.POST("/api/v1/prospect/upload", src.Parse)

	r.Run(":8000")
}
