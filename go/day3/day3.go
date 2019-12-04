package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"../lib"
)

func tests() {
	w11 := lib.ParseWire("R8,U5,L5,D3")
	if fmt.Sprintf("%s", w11) != "[R:8 U:5 L:5 D:3]" {
		fmt.Printf("test t1 failed! Expected [R:8 U:5 L:5 D:3] but got %s\n", w11)
	}

	w12 := lib.ParseWire("U7,R6,D4,L4")
	if fmt.Sprintf("%s", w12) != "[U:7 R:6 D:4 L:4]" {
		fmt.Printf("test t1 failed! Expected [U:7 R:6 D:4 L:4] but got %s\n", w12)
	}

	t1, _ := lib.GetWireOverlapDistance([]lib.Wire{w11, w12})
	if t1 != 6 {
		fmt.Printf("test t1 failed! Expected 6 but got %d\n", t1)
	} else {
		fmt.Println("test t1 passed!")
	}

	fmt.Println()

	w21 := lib.ParseWire("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	if fmt.Sprintf("%s", w21) != "[R:75 D:30 R:83 U:83 L:12 D:49 R:71 U:7 L:72]" {
		fmt.Printf("test t2 failed! Expected [R:75 D:30 R:83 U:83 L:12 D:49 R:71 U:7 L:72] but got %s\n", w21)
	}

	w22 := lib.ParseWire("U62,R66,U55,R34,D71,R55,D58,R83")
	if fmt.Sprintf("%s", w22) != "[U:62 R:66 U:55 R:34 D:71 R:55 D:58 R:83]" {
		fmt.Printf("test t2 failed! Expected [U:62 R:66 U:55 R:34 D:71 R:55 D:58 R:83] but got %s\n", w22)
	}

	t2, _ := lib.GetWireOverlapDistance([]lib.Wire{w21, w22})
	if t2 != 159 {
		fmt.Printf("test t2 failed! Expected 159 but got %d\n", t2)
	} else {
		fmt.Println("test t2 passed!")
	}

	fmt.Println()

	w31 := lib.ParseWire("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	w32 := lib.ParseWire("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	t3, _ := lib.GetWireOverlapDistance([]lib.Wire{w31, w32})

	if t3 != 135 {
		fmt.Printf("test t3 failed! Expected 135 but got %d\n", t3)
	} else {
		fmt.Println("test t3 passed!")
	}

	fmt.Println()

	w41 := lib.ParseWire("R8,U5,L5,D3")
	w42 := lib.ParseWire("U7,R6,D4,L4")
	_, t4 := lib.GetWireOverlapDistance([]lib.Wire{w41, w42})

	if t4 != 30 {
		fmt.Printf("test t4 failed! Expected 30 but got %d\n", t4)
	} else {
		fmt.Println("test t4 passed!")
	}

	w51 := lib.ParseWire("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	w52 := lib.ParseWire("U62,R66,U55,R34,D71,R55,D58,R83")
	_, t5 := lib.GetWireOverlapDistance([]lib.Wire{w51, w52})

	if t5 != 610 {
		fmt.Printf("test t5 failed! Expected 610 but got %d\n", t5)
	} else {
		fmt.Println("test t5 passed!")
	}

	w61 := lib.ParseWire("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	w62 := lib.ParseWire("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	_, t6 := lib.GetWireOverlapDistance([]lib.Wire{w61, w62})

	if t6 != 410 {
		fmt.Printf("test t6 failed! Expected 410 but got %d\n", t6)
	} else {
		fmt.Println("test t6 passed!")
	}

}

func part1() {
	raw, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	wireStrings := strings.Fields(string(raw))

	w1 := lib.ParseWire(wireStrings[0])
	w2 := lib.ParseWire(wireStrings[1])
	dist, _ := lib.GetWireOverlapDistance([]lib.Wire{w1, w2})

	fmt.Printf("Part1: %d\n", dist)

}

func part2() {

	raw, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	wireStrings := strings.Fields(string(raw))

	w1 := lib.ParseWire(wireStrings[0])
	w2 := lib.ParseWire(wireStrings[1])
	_, dist := lib.GetWireOverlapDistance([]lib.Wire{w1, w2})

	fmt.Printf("Part2: %d\n", dist)

}

func main() {
	fmt.Println("Day 3")
	fmt.Println()
	tests()
	part1()
	part2()
}
