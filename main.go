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
	src.ConnectDB()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// routes
	r.GET("/api/v1/prospect/", src.Index)
	// r.GET("/api/v1/prospect/upload", src.Parse) TOOD: add as route to template
	r.POST("/api/v1/prospect/score", src.Parsification)
	r.POST("/api/v1/prospect/suggest_score", src.Suggestion)

	r.Run(":8000")
}
