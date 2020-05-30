package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"../lib"
)

func tests() {

	fmt.Println("Test 11.")
	_, o11 := lib.RunIntcode([]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, []int{0})
	if o11[0] != 0 {
		fmt.Printf("Test 11 failed! Expected 0 but got: %d\n", o11)
	}
	fmt.Println()

	fmt.Println("Test 12.")
	_, o12 := lib.RunIntcode([]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, []int{1})
	if o12[0] != 1 {
		fmt.Printf("Test 12 failed! Expected 1 but got: %d\n", o12)
	}
	fmt.Println()

	fmt.Println("Test 13.")
	_, o13 := lib.RunIntcode([]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, []int{5})
	if o13[0] != 1 {
		fmt.Printf("Test 13 failed! Expected 1 but got: %d\n", o13)
	}
	fmt.Println()

	fmt.Println("Test 21.")
	_, o21 := lib.RunIntcode([]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, []int{0})
	if o21[0] != 0 {
		fmt.Printf("Test 21 failed! Expected 0 but got: %d\n", o21)
	}
	fmt.Println()

	fmt.Println("Test 22.")
	_, o22 := lib.RunIntcode([]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, []int{1})
	if o22[0] != 1 {
		fmt.Printf("Test 22 failed! Expected 1 but got: %d\n", o22)
	}
	fmt.Println()

	fmt.Println("Test 23.")
	_, o23 := lib.RunIntcode([]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, []int{5})
	if o23[0] != 1 {
		fmt.Printf("Test 23 failed! Expected 1 but got: %d\n", o23)
	}
	fmt.Println()

	fmt.Println("Test 31.")
	_, o31 := lib.RunIntcode([]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, []int{5})
	if o31[0] != 999 {
		fmt.Printf("Test 31 failed! Expected 999 but got: %d\n", o31)
	}
	fmt.Println()

	fmt.Println("Test 32.")
	_, o32 := lib.RunIntcode([]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, []int{7})
	if o32[0] != 999 {
		fmt.Printf("Test 32 failed! Expected 999 but got: %d\n", o32)
	}
	fmt.Println()

	fmt.Println("Test 33.")
	_, o33 := lib.RunIntcode([]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, []int{8})
	if o33[0] != 1000 {
		fmt.Printf("Test 33 failed! Expected 1000 but got: %d\n", o33)
	}
	fmt.Println()

	fmt.Println("Test 34.")
	_, o34 := lib.RunIntcode([]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, []int{9})
	if o34[0] != 1001 {
		fmt.Printf("Test 34 failed! Expected 1001 but got: %d\n", o34)
	}
	fmt.Println()

	fmt.Println("Test 35.")
	_, o35 := lib.RunIntcode([]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, []int{15})
	if o35[0] != 1001 {
		fmt.Printf("Test 35 failed! Expected 1001 but got: %d\n", o35)
	}
	fmt.Println()

	fmt.Println("Test 41.")
	_, o41 := lib.RunIntcode([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{7})
	if o41[0] != 0 {
		fmt.Printf("Test 41 failed! Expected 0 but got: %d\n", o41)
	}
	fmt.Println()

	fmt.Println("Test 42.")
	_, o42 := lib.RunIntcode([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{8})
	if o42[0] != 1 {
		fmt.Printf("Test 42 failed! Expected 1 but got: %d\n", o42)
	}
	fmt.Println()

	fmt.Println("Test 43.")
	_, o43 := lib.RunIntcode([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{9})
	if o43[0] != 0 {
		fmt.Printf("Test 43 failed! Expected 0 but got: %d\n", o43)
	}
	fmt.Println()

	fmt.Println("Test 51.")
	_, o51 := lib.RunIntcode([]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{7})
	if o51[0] != 1 {
		fmt.Printf("Test 51 failed! Expected 1 but got: %d\n", o51)
	}
	fmt.Println()

	fmt.Println("Test 52.")
	_, o52 := lib.RunIntcode([]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{8})
	if o52[0] != 0 {
		fmt.Printf("Test 52 failed! Expected 0 but got: %d\n", o52)
	}
	fmt.Println()

	fmt.Println("Test 53.")
	_, o53 := lib.RunIntcode([]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{9})
	if o53[0] != 0 {
		fmt.Printf("Test 53 failed! Expected 0 but got: %d\n", o53)
	}
	fmt.Println()

	fmt.Println("Test 61.")
	_, o61 := lib.RunIntcode([]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []int{7})
	if o61[0] != 0 {
		fmt.Printf("Test 61 failed! Expected 0 but got: %d\n", o61)
	}
	fmt.Println()

	fmt.Println("Test 62.")
	_, o62 := lib.RunIntcode([]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []int{8})
	if o62[0] != 1 {
		fmt.Printf("Test 62 failed! Expected 1 but got: %d\n", o62)
	}
	fmt.Println()

	fmt.Println("Test 63.")
	_, o63 := lib.RunIntcode([]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []int{9})
	if o63[0] != 0 {
		fmt.Printf("Test 63 failed! Expected 0 but got: %d\n", o63)
	}
	fmt.Println()

	fmt.Println("Test 71.")
	_, o71 := lib.RunIntcode([]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int{7})
	if o71[0] != 1 {
		fmt.Printf("Test 71 failed! Expected 1 but got: %d\n", o71)
	}
	fmt.Println()

	fmt.Println("Test 72.")
	_, o72 := lib.RunIntcode([]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int{8})
	if o72[0] != 0 {
		fmt.Printf("Test 72 failed! Expected 0 but got: %d\n", o72)
	}
	fmt.Println()

	fmt.Println("Test 73.")
	_, o73 := lib.RunIntcode([]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int{9})
	if o73[0] != 0 {
		fmt.Printf("Test 73 failed! Expected 0 but got: %d\n", o73)
	}
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

	data, _ = lib.RunIntcode(data, []int{1})

	fmt.Printf("Part 1 should be: 4511442\n")
}

func part2() {
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

	data, _ = lib.RunIntcode(data, []int{5})

	fmt.Printf("Part 2 should be: 12648139\n")
}

func main() {
	fmt.Println("Day 5")
	fmt.Println()
	tests()
	part1()
	part2()
}
