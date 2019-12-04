package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"../lib"
)

func part2() {

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

		// for each module mass, calculate its fuel and add it to the total. Then, treat the fuel amount you just calculated as the input mass and repeat the process, continuing until a fuel requirement is zero or negative.

		for {
			fuel := lib.FuelForMass(mass)
			if fuel > 0 {
				total += fuel
				mass = fuel
			} else {
				break
			}
		}
	}

	fmt.Printf("Part2 total: %d\n", total)
}
