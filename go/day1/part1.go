package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"../lib"
)

func part1() {

	dat, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	modules := strings.Fields(string(dat))

	var total int64

	for _, m := range modules {
		mass, err := strconv.ParseInt(m, 10, 64)
		if err != nil {
			panic(err)
		}

		fuel := lib.FuelForMass(mass)

		// fmt.Printf("%d => %d\n", mass, fuel)
		total += fuel
	}

	fmt.Printf("Part 1 total: %d\n", total)
}

func main() {
	part1()
	part2()
}
