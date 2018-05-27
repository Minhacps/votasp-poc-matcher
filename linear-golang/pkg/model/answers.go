package model

import (
	"encoding/json"
		)

//Answers represents a set of answers from an individual
type Answers map[string][]int8

func (t *Answers) UnmarshalJSON(b []byte) error {
	var answersAsString map[string][]string

	err := json.Unmarshal(b, &answersAsString)
	if err != nil {
		return err
	}

	toInt := map[string]int8{
		"CP": 2,
		"C": 1,
		"I": 0,
		"D": -1,
		"DP": -2,
	}

	*t = make(map[string][]int8)

	for id,answers := range answersAsString {
		intAnswers := make([]int8, len(answers))
		for i,answer := range answers {
			intAnswers[i] = toInt[answer]
		}
		(*t)[id] = intAnswers
	}

	return nil
}
