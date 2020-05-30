package lib

import (
	"fmt"
	"strconv"
)

// Days:
// - 2
// - 5

// RunIntcode processes an intcode
func RunIntcode(intcode []int, input []int) ([]int, []int) {

	var (
		ip     int
		s      int
		run    bool = true
		ii     int
		output []int
	)

	for run {
		op, mode1, mode2, mode3 := splitOpcode(intcode[ip])

		switch op {
		case 1: // Addition
			s = 4
			p1 := getParameter(&intcode, ip, 1, mode1)
			p2 := getParameter(&intcode, ip, 2, mode2)
			setParameter(&intcode, ip, 3, mode3, p1+p2)
		case 2: // Multiplication
			s = 4
			p1 := getParameter(&intcode, ip, 1, mode1)
			p2 := getParameter(&intcode, ip, 2, mode2)
			setParameter(&intcode, ip, 3, mode3, p1*p2)
		case 3: // Input
			s = 2
			var t int
			if ii < len(input) {
				t = input[ii]
				ii++
			} else {
				fmt.Printf("Input: ")
				fmt.Scanf("%d", &t)
			}
			setParameter(&intcode, ip, 1, mode1, t)
		case 4: // Output
			s = 2
			p1 := getParameter(&intcode, ip, 1, mode1)
			// fmt.Printf("Output: %d\n", p1)
			output = append(output, p1)
		case 5: // Jump if True
			s = 3
			p1 := getParameter(&intcode, ip, 1, mode1)
			p2 := getParameter(&intcode, ip, 2, mode2)
			if p1 != 0 {
				ip = p2
				s = 0
			}
		case 6: // Jump if False
			s = 3
			p1 := getParameter(&intcode, ip, 1, mode1)
			p2 := getParameter(&intcode, ip, 2, mode2)
			if p1 == 0 {
				ip = p2
				s = 0
			}
		case 7: // Less than
			s = 4
			p1 := getParameter(&intcode, ip, 1, mode1)
			p2 := getParameter(&intcode, ip, 2, mode2)
			if p1 < p2 {
				setParameter(&intcode, ip, 3, mode3, 1)
			} else {
				setParameter(&intcode, ip, 3, mode3, 0)
			}
		case 8: // Equals
			s = 4
			p1 := getParameter(&intcode, ip, 1, mode1)
			p2 := getParameter(&intcode, ip, 2, mode2)
			if p1 == p2 {
				setParameter(&intcode, ip, 3, mode3, 1)
			} else {
				setParameter(&intcode, ip, 3, mode3, 0)
			}
		case 99:
			s = 0
			run = false
		}

		ip += s
	}

	return intcode, output
}

func splitOpcode(n int) (int, int, int, int) {
	padded := fmt.Sprintf("%d", n)
	for len(padded) < 5 {
		padded = "0" + padded
	}

	op, _ := strconv.ParseInt(padded[3:5], 10, 32)
	mode1, _ := strconv.ParseInt(padded[2:3], 10, 32)
	mode2, _ := strconv.ParseInt(padded[1:2], 10, 32)
	mode3, _ := strconv.ParseInt(padded[0:1], 10, 32)

	return int(op), int(mode1), int(mode2), int(mode3)
}

func getParameter(intcode *[]int, ip int, off int, mode int) int {
	switch mode {
	case 0:
		return (*intcode)[(*intcode)[ip+off]]
	case 1:
		return (*intcode)[ip+off]
	default:
		fmt.Printf("getParameter given unknown mode: %d\n", mode)
		return 0
	}
}

func setParameter(intcode *[]int, ip int, off int, mode int, val int) {
	switch mode {
	case 0:
		(*intcode)[(*intcode)[ip+off]] = val
	}
}
