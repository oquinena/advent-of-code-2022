package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type stack struct {
	box []rune
}

func (s *stack) push(r []rune) {
	s.box = append(s.box, r...)
}

func (s *stack) pop(n int) []rune {
	if len(s.box) == 0 {
		return nil
	}
	r := s.box[len(s.box)-n : len(s.box)]
	s.box = s.box[:len(s.box)-n]
	return r
}

func (s *stack) addToBottom(r rune) {
	s.box = append([]rune{r}, s.box...)
}

func simulateMovement(cratesToMove int) string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	stacks := make([]stack, 9)
	var tops string

	scanner.Scan()
	for scanner.Text() != " 1   2   3   4   5   6   7   8   9 " {
		for i, r := range scanner.Text() {
			if r != ' ' && r != '[' && r != ']' {
				stacks[i/4].addToBottom(r)
			}
		}
		scanner.Scan()
	}
	scanner.Scan()
	for scanner.Scan() {
		var movement, from, to int
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &movement, &from, &to)

		if cratesToMove <= 1 {
			for move := 0; move < movement; move++ {
				stacks[to-1].push(stacks[from-1].pop(cratesToMove))
			}
		} else {
			stacks[to-1].push(stacks[from-1].pop(movement))
		}
	}
	for _, s := range stacks {
		tops = tops + string(s.pop(1))
	}
	return tops
}

func main() {

	part1 := simulateMovement(1)
	part2 := simulateMovement(3)

	fmt.Printf("(PART 1) Crates on top: %s\n(PART 2) Crates on top: %s\n", part1, part2)
}
