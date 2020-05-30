package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"../lib"
)

func tests() {
	fmt.Println("Test 1")
	signal := runAmp([]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, []int{4, 3, 2, 1, 0})
	fmt.Printf("Signal: %d\n", signal)
	fmt.Println("Should be: 43210")
	fmt.Println()

	fmt.Println("Test 2")
	signal = runAmp([]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
		101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}, []int{0, 1, 2, 3, 4})
	fmt.Printf("Signal: %d\n", signal)
	fmt.Println("Should be: 54321")
	fmt.Println()

	fmt.Println("Test 3")
	signal = runAmp([]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
		1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}, []int{1, 0, 4, 3, 2})
	fmt.Printf("Signal: %d\n", signal)
	fmt.Println("Should be: 65210")
	fmt.Println()

}

func runAmp(intcode []int, settings []int) int {
	tmp := make([]int, len(intcode))
	signal := 0

	for i := 0; i < 5; i++ {
		copy(tmp, intcode)
		_, o := lib.RunIntcode(tmp, []int{settings[i], signal})
		signal = o[0]
	}

	return signal
}

func part1() {
	fmt.Println("Part 1.")

	raw, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	trim := strings.TrimSpace(string(raw))

	split := strings.Split(trim, ",")
	intcode := make([]int, len(split))
	for i := range intcode {
		d, _ := strconv.ParseInt(split[i], 10, 32)
		intcode[i] = int(d)
	}

	max := 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == i {
				continue
			}
			for k := 0; k < 5; k++ {
				if k == i || k == j {
					continue
				}
				for l := 0; l < 5; l++ {
					if l == i || l == j || l == k {
						continue
					}
					for m := 0; m < 5; m++ {
						if m == i || m == j || m == k || m == l {
							continue
						}
						signal := runAmp(intcode, []int{i, j, k, l, m})
						if signal > max {
							max = signal
							fmt.Printf("%d %d %d %d %d: %d\n", i, j, k, l, m, signal)
						}

					}
				}
			}
		}
	}

	fmt.Printf("Max: %d\n", max)
}

func main() {
	fmt.Println("Day 7")
	fmt.Println()

	tests()
	fmt.Println()

	part1()
	fmt.Println()
}
