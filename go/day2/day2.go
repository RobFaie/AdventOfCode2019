package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"../lib"
)

func tests() {
	t1, _ := lib.RunIntcode([]int{1, 0, 0, 3, 99}, nil)
	// fmt.Printf("%s", t1)
	if t1[3] != 2 {
		fmt.Printf("test t1 failed! Expected 2 but got %d\n", t1[3])
	}

	t2, _ := lib.RunIntcode([]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, nil)
	if t2[0] != 3500 {
		fmt.Printf("test t2 failed! Expected 3500 but got %d\n", t2[0])
	}

	t3, _ := lib.RunIntcode([]int{1, 0, 0, 0, 99}, nil)
	if t3[0] != 2 {
		fmt.Printf("test t3 failed! Expected 2 but got %d\n", t3[0])
	}

	t4, _ := lib.RunIntcode([]int{2, 3, 0, 3, 99}, nil)
	if t4[3] != 6 {
		fmt.Printf("test t4 failed! Expected 6 but got %d\n", t4[3])
	}

	t5, _ := lib.RunIntcode([]int{2, 4, 4, 5, 99, 0}, nil)
	if t5[5] != 9801 {
		fmt.Printf("test t5 failed! Expected 9801 but got %d\n", t5[5])
	}

	t6, _ := lib.RunIntcode([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, nil)
	if t6[0] != 30 {
		fmt.Printf("test t6 failed! Expected 30 but got %d\n", t6[0])
	}

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

	data[1] = 12
	data[2] = 2

	data, _ = lib.RunIntcode(data, nil)

	fmt.Printf("Part1: %d\n", data[0])
	fmt.Println("Should have been: 5434663")
}

func part2() {
	raw, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	strs := strings.Split(string(raw), ",")

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {

			data := make([]int, len(strs))
			for i := range data {
				d, _ := strconv.ParseInt(strs[i], 10, 32)
				data[i] = int(d)
			}

			data[1] = int(i)
			data[2] = int(j)

			data, _ = lib.RunIntcode(data, nil)

			if data[0] == 19690720 {
				fmt.Printf("Part2: %d\n", 100*i+j)
				fmt.Println("Should have been: 4559")
				return
			}
		}
	}

}

func main() {
	fmt.Println("Day 2")
	fmt.Println()
	tests()
	part1()
	fmt.Println()
	part2()
}
