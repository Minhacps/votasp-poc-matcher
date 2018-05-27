package main

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"./pkg/model"
	"time"
	"log"
	"./pkg/service"
	"sort"
	)

func main() {
	rawCandidatesAnswers, err := ioutil.ReadFile(os.Args[1]); if err != nil {
		panic(err)
	}

	rawElectorsAnswers, err := ioutil.ReadFile(os.Args[2]); if err != nil {
		panic(err)
	}
    var candidatesAnswers model.Answers
    json.Unmarshal(rawCandidatesAnswers, &candidatesAnswers)

	var electorsAnswers model.Answers
	json.Unmarshal(rawElectorsAnswers, &electorsAnswers)

	start := time.Now()

	for _,electorAnswers := range electorsAnswers {
		candidatesScores := make(model.MatchScoreList, len(candidatesAnswers))
		for candidateId,candidateAnswers := range candidatesAnswers {
			score := service.Score(candidateAnswers, electorAnswers)
			candidatesScores = append(candidatesScores, model.MatchScore{candidateId, score})
		}
		sort.Sort(candidatesScores)
		//topTen := candidatesScores[len(candidatesScores)-10:]
	}

	elapsed := time.Since(start)
	log.Printf("Took %s to match %d elector answers with %d candidate answers", elapsed,len(electorsAnswers), len(candidatesAnswers))

}

