package main

import "fmt"

func part1() {
	min := 372304
	max := 847060

	total := 0

	for i := min; i <= max; i++ {
		s := fmt.Sprintf("%d", i)
		// fmt.Printf("%d => %s\n", i, s)
		if s[0] <= s[1] && s[1] <= s[2] && s[2] <= s[3] && s[3] <= s[4] && s[4] <= s[5] {
			if s[0] == s[1] || s[1] == s[2] || s[2] == s[3] || s[3] == s[4] || s[4] == s[5] {
				total++
			}
		}
	}

	fmt.Printf("Part1: %d\n", total)
}

func part2() {
	min := 372304
	max := 847060

	total := 0

	for i := min; i <= max; i++ {
		s := fmt.Sprintf("%d", i)
		// fmt.Printf("%d => %s\n", i, s)
		if s[0] <= s[1] && s[1] <= s[2] && s[2] <= s[3] && s[3] <= s[4] && s[4] <= s[5] {
			if (s[0] == s[1] && s[1] != s[2]) ||
				(s[1] == s[2] && s[0] != s[1] && s[2] != s[3]) ||
				(s[2] == s[3] && s[1] != s[2] && s[3] != s[4]) ||
				(s[3] == s[4] && s[2] != s[3] && s[4] != s[5]) ||
				(s[4] == s[5] && s[3] != s[4]) {
				total++
				fmt.Println(s)
			}
		}
	}

	fmt.Printf("Part2: %d\n", total)
}

func main() {
	part1()
	part2()
}
