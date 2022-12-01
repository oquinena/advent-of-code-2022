package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calories(calories string) (int, int) {
	var (
		elfCalories []int
		counter     = 0
		most        = 0
		total       = 0
	)

	caloryLines := strings.Split(calories, "\n")

	for _, line := range caloryLines {
		if line != "" {
			i, _ := strconv.Atoi(line)
			counter += i
		} else {
			elfCalories = append(elfCalories, counter)
			counter = 0
		}
	}

	sort.Ints(elfCalories)
	for _, i := range elfCalories[len(elfCalories)-3:] {
		most = i
		total += i
	}
	return most, total
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	most, total := calories(string(input))
	fmt.Printf("(PART 1) Most calories: %v\n(PART 2) Total calories: %v\n", most, total)
}
