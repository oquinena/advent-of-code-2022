package main

import (
	"bufio"
	"fmt"
	"os"
)

func compareRanges(f *os.File) (int, int) {
	scanner := bufio.NewScanner(f)
	part1, part2 := 0, 0
	for scanner.Scan() {
		var p1, p2 struct{ l, r int }
		_, err := fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &p1.l, &p1.r, &p2.l, &p2.r)
		if err != nil {
			panic(err)
		}
		if p1.l <= p2.l && p1.r >= p2.r || p2.l <= p1.l && p2.r >= p1.r {
			part1 += 1
		}
		if p1.r >= p2.l && p2.r >= p1.l {
			part2 += 1
		}
	}
	return part1, part2
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	part1, part2 := compareRanges(f)
	fmt.Printf("(PART1) Duplicate assignments: %v\n(PART2) Overlapping assignments: %v\n", part1, part2)
}
