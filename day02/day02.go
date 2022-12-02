package main

import (
	"fmt"
	"os"
	"strings"
)

func part1(r string) int {
	line := strings.Split(r, "\n")
	score := 0

	possibleScores := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	for _, v := range line {
		score += possibleScores[v]
	}

	return score
}

func part2(r string) int {
	line := strings.Split(r, "\n")
	score := 0

	possibleScores := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}

	for _, v := range line {
		score += possibleScores[v]
	}

	return score
}

func main() {

	f, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	score1 := part1(string(f))
	score2 := part2(string(f))

	fmt.Printf("(PART1) Total score: %v\n(PART2) Total score: %v\n", score1, score2)
}
