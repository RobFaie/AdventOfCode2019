package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Day 6.")
	fmt.Println()

	dat, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	orbits := strings.Fields(string(dat))

	m := make(map[string]string)

	for _, v := range orbits {
		// fmt.Printf("Orbit %d\n", i)
		parts := strings.Split(v, ")")
		m[parts[1]] = parts[0]
	}

	total := 0

	// fmt.Printf("%s\n", m)
	for k := range m {
		for k != "COM" {
			total++
			k = m[k]
		}
	}

	fmt.Printf("Orbits: %d\n", total)
	fmt.Println()

	y := "YOU"
	var ys []string
	for y != "COM" {
		ys = append([]string{m[y]}, ys...)
		y = m[y]
	}

	// fmt.Printf("YOU => %s\n", ys)

	s := "SAN"
	var ss []string
	for s != "COM" {
		ss = append([]string{m[s]}, ss...)
		s = m[s]
	}

	// fmt.Printf("SAN => %s\n", ss)

	i := 0
	for ys[i] == ss[i] {
		i++
	}

	// fmt.Printf("YOU => %s\n", ys[i:])
	// fmt.Printf("SAN => %s\n", ss[i:])

	fmt.Printf("Jump distance: %d", len(ys)+len(ss)-i-i)

	fmt.Println()
}
