package service

func Score(candidateAnswers []int8, electorAnswers []int8)  int{
	var score int
	for i := 0; i < len(candidateAnswers); i++ {
		if electorAnswers[i] == 0 {
			continue
		}
		if candidateAnswers[i] == electorAnswers[i] {
			score += 2
			continue
		}
		if candidateAnswers[i] * electorAnswers[i] > 0 {
			score += 1
			continue
		}
	}

	return score
}
