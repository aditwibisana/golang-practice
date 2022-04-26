package main

import "fmt"

func main() {
	/*
		Score Rank:
		90-100: A
		80-89: B
		70-79: C
		60-69: D
		0-59: E
	*/
	res := ScoreRank(50)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := ScoreRankCorrect(arr)
	// fmt.Println(resCorrect)
	res = ScoreRankCorrect(100)
	fmt.Println(res)
}

func ScoreRank(score int) string {
	if score > 100 || score < 0 {
		return "INVALID"
	}
	var rank string
	switch {
	case score > 60:
		rank = "D"
	case score > 70:
		rank = "C"
	case score > 80:
		rank = "B"
	case score > 90:
		rank = "A"
	default:
		rank = "E"
	}

	return rank
}

func ScoreRankCorrect(score int) string {
	if score > 100 || score < 0 {
		return "INVALID"
	}
	var rank string
	if score == 100 {
		rank = "A"
	} else if score >= 90 {
		rank = "A"
	} else if score >= 80 {
		rank = "B"
	} else if score >= 70 {
		rank = "C"
	} else if score >= 60 {
		rank = "D"
	} else {
		rank = "E"
	}
	return rank // TODO: replace this
}