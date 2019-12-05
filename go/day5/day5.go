package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"../lib"
)

func tests() {

	fmt.Println("Test 1.")
	lib.RunIntcode([]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9})
	fmt.Println()

	fmt.Println("Test 2.")
	lib.RunIntcode([]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1})
	fmt.Println()

	fmt.Println("Test 3.")
	lib.RunIntcode([]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99})
	fmt.Println()

	fmt.Println("Test 4.")
	lib.RunIntcode([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8})
	fmt.Println()

	fmt.Println("Test 5.")
	lib.RunIntcode([]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8})
	fmt.Println()

	fmt.Println("Test 6.")
	lib.RunIntcode([]int{3, 3, 1108, -1, 8, 3, 4, 3, 99})
	fmt.Println()

	fmt.Println("Test 7.")
	lib.RunIntcode([]int{3, 3, 1107, -1, 8, 3, 4, 3, 99})
	fmt.Println()
}

func part1() {
	raw, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(raw), ",")
	data := make([]int, len(strs))
	for i := range data {
		d, _ := strconv.ParseInt(strs[i], 10, 32)
		data[i] = int(d)
	}

	data = lib.RunIntcode(data)

	fmt.Printf("Part 1 complete.")
}

func part2() {
	raw, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 2: %s\n", raw)
}

func main() {
	fmt.Println("Day 5")
	fmt.Println()
	// tests()
	part1()
	// part2()
}
