package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"../lib"
)

func tests() {
	t1 := lib.RunIntcode([]int64{1, 0, 0, 3, 99})
	// fmt.Printf("%s", t1)
	if t1[3] != 2 {
		fmt.Printf("test t1 failed! Expected 2 but got %d\n", t1[3])
	}

	t2 := lib.RunIntcode([]int64{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50})
	if t2[0] != 3500 {
		fmt.Printf("test t2 failed! Expected 3500 but got %d\n", t2[0])
	}

	t3 := lib.RunIntcode([]int64{1, 0, 0, 0, 99})
	if t3[0] != 2 {
		fmt.Printf("test t3 failed! Expected 2 but got %d\n", t3[0])
	}

	t4 := lib.RunIntcode([]int64{2, 3, 0, 3, 99})
	if t4[3] != 6 {
		fmt.Printf("test t4 failed! Expected 6 but got %d\n", t4[3])
	}

	t5 := lib.RunIntcode([]int64{2, 4, 4, 5, 99, 0})
	if t5[5] != 9801 {
		fmt.Printf("test t5 failed! Expected 9801 but got %d\n", t5[5])
	}

	t6 := lib.RunIntcode([]int64{1, 1, 1, 4, 99, 5, 6, 0, 99})
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
	data := make([]int64, len(strs))
	for i := range data {
		data[i], _ = strconv.ParseInt(strs[i], 10, 64)
	}

	data[1] = 12
	data[2] = 2

	data = lib.RunIntcode(data)

	fmt.Printf("Part1: %d\n", data[0])
}

func part2() {
	raw, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	strs := strings.Split(string(raw), ",")

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {

			data := make([]int64, len(strs))
			for i := range data {
				data[i], _ = strconv.ParseInt(strs[i], 10, 64)
			}

			data[1] = int64(i)
			data[2] = int64(j)

			data = lib.RunIntcode(data)

			if data[0] == 19690720 {
				fmt.Printf("Part2: %d", 100*i+j)
				return
			}
		}
	}

}

func main() {
	fmt.Println("Day 2")
	tests()
	part1()
	part2()
}
