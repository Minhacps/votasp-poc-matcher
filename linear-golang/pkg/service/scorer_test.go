package service_test

import (
		"testing"
				. "../service"
)

var scoreTests = []struct {
	candidateAnswers  []int8
	electorAnswers  []int8
	expectedScore int
}{
	{[]int8{2}, []int8{2}, 2},
	{[]int8{1}, []int8{2}, 1},
	{[]int8{-1}, []int8{2}, 0},
	{[]int8{-2}, []int8{2}, 0},
	{[]int8{2}, []int8{1}, 1},
	{[]int8{1}, []int8{1}, 2},
	{[]int8{-1}, []int8{1}, 0},
	{[]int8{-2}, []int8{1}, 0},
	{[]int8{2}, []int8{0}, 0},
	{[]int8{1}, []int8{0}, 0},
	{[]int8{-1}, []int8{0}, 0},
	{[]int8{-2}, []int8{0}, 0},
	{[]int8{2}, []int8{-1}, 0},
	{[]int8{1}, []int8{-1}, 0},
	{[]int8{-1}, []int8{-1}, 2},
	{[]int8{-2}, []int8{-1}, 1},
	{[]int8{2}, []int8{-2}, 0},
	{[]int8{1}, []int8{-2}, 0},
	{[]int8{-1}, []int8{-2}, 1},
	{[]int8{-2}, []int8{-2}, 2},
}

func TestScoring(t *testing.T)  {
	for _,testData := range scoreTests {


		expectedScore := testData.expectedScore
		actualScore := Score(testData.candidateAnswers, testData.electorAnswers)

		if actualScore != expectedScore {
			t.Errorf("Expected score %d for candidate %q and elector %q, got %d", expectedScore, testData.candidateAnswers, testData.electorAnswers, actualScore)
		}
	}

}
