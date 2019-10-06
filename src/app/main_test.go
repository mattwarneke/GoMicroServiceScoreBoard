package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_findScoresForUser(t *testing.T) {
	score1 := Score{15, "matt", time.Now(), "level", "1.2.3"}
	score2 := Score{10, "bob", time.Now(), "level", "1.2.3"}
	score3 := Score{20, "matt", time.Now(), "level", "1.2.3"}
	var scoresEmpty []Score

	tests := []struct {
		name string
		user string
		want []Score
	}{
		{"matt score", "matt", []Score{score1, score3}},
		{"bob score", "bob", []Score{score2}},
		{"ted empty score", "ted", scoresEmpty},
	}

	AddScore(score1)
	AddScore(score2)
	AddScore(score3)

	for _, tt := range tests {
		res := findScoresForUser(tt.user)
		if reflect.DeepEqual(res, tt.want) {
			t.Log(tt.name + " Passed")
		} else {
			t.Errorf(tt.name+" Failed got %v but expected %v", res, tt.want)
		}
	}
}
