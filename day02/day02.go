package main

import (
	"fmt"
	"os"
	"strings"
)

func scores(r string) (int, int) {
	lines := strings.Split(r, "\n")
	score1, score2 := 0, 0

	possibleScores := map[string]struct{ a, b int }{
		"A X": {4, 3},
		"A Y": {8, 4},
		"A Z": {3, 8},
		"B X": {1, 1},
		"B Y": {5, 5},
		"B Z": {9, 9},
		"C X": {7, 2},
		"C Y": {2, 6},
		"C Z": {6, 7},
	}

	for _, l := range lines {
		score1 += possibleScores[l].a
		score2 += possibleScores[l].b
	}

	return score1, score2
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	score1, score2 := scores(string(f))

	fmt.Printf("(PART1) Total score: %v\n(PART2) Total score: %v\n", score1, score2)
}
