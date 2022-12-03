package main

import (
	"fmt"
	"os"
	"strings"
)

func rucksacks(r string) (int, int) {
	rucksack := strings.Split(r, "\n")
	groups := make([]string, 0)
	part1, part2 := 0, 0

	priority := map[rune]int{
		'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6,
		'g': 7, 'h': 8, 'i': 9, 'j': 10, 'k': 11,
		'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16,
		'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21,
		'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26,
		'A': 27, 'B': 28, 'C': 29, 'D': 30, 'E': 31,
		'F': 32, 'G': 33, 'H': 34, 'I': 35, 'J': 36,
		'K': 37, 'L': 38, 'M': 39, 'N': 40, 'O': 41,
		'P': 42, 'Q': 43, 'R': 44, 'S': 45, 'T': 46,
		'U': 47, 'V': 48, 'W': 49, 'X': 50, 'Y': 51,
		'Z': 52,
	}

	for _, r := range rucksack {
		first, second := r[:len(r)/2], r[len(r)/2:]
		for _, a := range first {
			if strings.ContainsRune(second, a) {
				fmt.Println(priority[a])
				part1 += priority[a]
				break
			}
		}
		groups = append(groups, r)
		if len(groups) == 3 {
			for _, b := range groups[0] {
				if strings.ContainsRune(groups[1], b) && strings.ContainsRune(groups[2], b) {
					part2 += priority[b]
					break
				}
			}
			groups = make([]string, 0)
		}
	}
	return part1, part2
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	part1, part2 := rucksacks(string(f))
	fmt.Printf("(PART1) Sum of priorities: %v\n(PART2) Sum of groups: %v\n", part1, part2)
}
