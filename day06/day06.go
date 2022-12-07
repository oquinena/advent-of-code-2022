package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func identifier(code int) int {
	var a = 0
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()
	for i := range s.Text() {
		c := make(map[byte]int)
		for l := 0; l < code; l++ {
			c[s.Text()[i+l]] = 1
		}
		if len(c) == code {
			a = i + code
			break
		}
	}
	return a
}

func main() {
	part1 := identifier(4)
	part2 := identifier(14)
	fmt.Printf("(PART 1) Number of digits before code: %d\n(PART 2) Number of digits before code: %d\n", part1, part2)
}
