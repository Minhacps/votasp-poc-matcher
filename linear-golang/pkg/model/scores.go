package model

type MatchScore struct {
	CandidateId string
	Score       int
}

type MatchScoreList []MatchScore

func (s MatchScoreList) Len() int { return len(s) }
func (s MatchScoreList) Less(i, j int) bool { return s[i].Score < s[j].Score }
func (s MatchScoreList) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
