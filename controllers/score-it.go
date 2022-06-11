package controllers

import (
	"fmt"
	"log"
	"os"

	gp "github.com/exsocial/goperspective"
)

func AnalyzeIt(payload string) {
	client := gp.NewClient(os.Getenv("API_KEY"))

	data := gp.AnalyzeRequest{
		Comment: gp.AnalyzeRequestComment{
			Text: payload,
		},
		ReqAttr: map[gp.Attribute]gp.AnalyzeRequestAttr{
			gp.Toxicity: {},
			gp.SevereToxicity: {},
			gp.Threat:         {},
			gp.IdentityAttack: {},
		},
	}

	r, err := client.AnalyzeComment(data)
	if err != nil {
		log.Fatal(err)
	}

	//Retrieving values
	for name, as := range r.AttributeScores {
		fmt.Println(name, as.SummaryScore.Value)
	}
}
