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
			gp.Toxicity:       {},
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

func SuggestScore(payload string) {
	/*
		lets you give the API feedback by allowing you to suggest
		a score that you think the API should have returned.
		You can use this method if you disagree with a score and
		would like to improve the attribute. All submissions to
		SuggestCommentScore are stored and may be used to improve
		the API and related services.
	*/
	client := gp.NewClient(os.Getenv("API_KEY"))

	data := gp.SuggestRequest{
		Comment: gp.AnalyzeRequestComment{
			Text: payload,
		},
		AttributeScores: map[gp.Attribute]gp.AttributeScore{
			gp.Toxicity: {},
			gp.Threat:   {},
		},
	}

	// TODO: update above
	// data := gp.SuggestRequest{
	// 	Comment: gp.AnalyzeRequestComment{
	// 		Text: "You are fucking stupid, aren't you?",
	// 	},
	// 	AttributeScores: map[gp.Attribute]gp.AttributeScore{
	// 		gp.Toxicity: {
	// 			SummaryScore: gp.Score{
	// 				Value: 1.0,
	// 			},
	// 		},
	// 		gp.Threat: {
	// 			SummaryScore: gp.Score{
	// 				Value: 0,
	// 			},
	// 		},
	// 	},
	// }

	_, err := client.SuggestCommentScore(data)
	//If there is no error, then suggestion is accepted.
	if err != nil {
		log.Fatal(err)
	}
}
