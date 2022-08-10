package src

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/cloudguruab/prospect/controllers"

	//    "io/ioutil"
	//    "github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin"
)

/*
	===================
	Serializers
	===================
	* Home serializer
	* Login serializer
	* Score Summary serializer
	* Parsification/Scoring serializer
	* Score suggestion serializer
*/
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Score summary page"})
}

func LoginPage(c *gin.Context) {}

// for uploading files, parsification/score suggestions
// for running action actions (remove route logic from
// parse/suggest - hold as functions that work in this route here)
func Upload(c *gin.Context) {}

// TODO: need to update how failure is handled and logging
// write method for triggering scores
// perspective api res loaded into
// tmp schema.go and saved to db
func Parsification(c *gin.Context) {
	ok, err := CheckEmpty("./tmp/")
	var score_bucket []string

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	}

	if ok == false {
		data, err := os.Open("./tmp/payload.txt")
		Check(err)

		defer data.Close()

		obj := bufio.NewScanner(data)
		for obj.Scan() {
			score_bucket = append(score_bucket, obj.Text())
		}

		Check(obj.Err())

		c.JSON(http.StatusOK, gin.H{"response": "Parsing..."})
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "Directory empty.."})
	}

	// taking this further we can create struct - return struct over api
	for i := 0; i <= len(score_bucket); i++ {
		fmt.Printf("%s \n", score_bucket[i])
		controllers.AnalyzeIt(score_bucket[i])
	}

	// to continue...
	// trigger cloud function || trigger controller node
	// we can either invoke cloud function from event change in cloud store
	// or we can invoke parsification directly from api call using controller
}

// FIXME: update once template implemented
// will be for giving ai model feedback on returned
// scores so we can consistently enhance our UX.
func Suggestion(c *gin.Context) {
	ok, err := CheckEmpty("./tmp/")
	var score_bucket []string

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	}

	if ok == false {
		data, err := os.Open("./tmp/payload.txt")
		Check(err)

		defer data.Close()

		obj := bufio.NewScanner(data)
		for obj.Scan() {
			score_bucket = append(score_bucket, obj.Text())
		}

		Check(obj.Err())

		c.JSON(http.StatusOK, gin.H{"response": "Parsing..."})
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "Directory empty.."})
	}

	for i := 0; i <= len(score_bucket); i++ {
		fmt.Printf("%s \n", score_bucket[i])
		controllers.SuggestScore(score_bucket[i])
	}
}
