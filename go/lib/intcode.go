package lib

import (
	"fmt"
	"strconv"
)

// RunIntcode processes an intcode
func RunIntcode(intcode []int) []int {

	var (
		ip  int
		s   int
		run bool = true
		// mode int
	)

	for run {
		op, mode1, mode2, mode3 := splitOpcode(intcode[ip])

		// fmt.Printf("IP: %d OP: %d\n", ip, op)

		switch op {
		case 1:
			{ // Addition
				s = 4
				p1 := getParameter(&intcode, ip, 1, mode1)
				p2 := getParameter(&intcode, ip, 2, mode2)
				setParameter(&intcode, ip, 3, mode3, p1+p2)
			}
		case 2:
			{ // Multiplication
				s = 4
				p1 := getParameter(&intcode, ip, 1, mode1)
				p2 := getParameter(&intcode, ip, 2, mode2)
				setParameter(&intcode, ip, 3, mode3, p1*p2)
			}

		case 3:
			{ // Input
				s = 2
				var t int
				fmt.Printf("Input: ")
				fmt.Scanf("%d", &t)
				setParameter(&intcode, ip, 1, mode1, t)
			}
		case 4:
			{ // Output
				s = 2
				p1 := getParameter(&intcode, ip, 1, mode1)
				fmt.Printf("Output: %d\n", p1)
			}
		case 5:
			{ // Jump if True
				s = 3
				p1 := getParameter(&intcode, ip, 1, mode1)
				p2 := getParameter(&intcode, ip, 2, mode2)
				if p1 != 0 {
					ip = p2
					s = 0
				}
			}
		case 6:
			{ // Jump if False
				s = 3
				p1 := getParameter(&intcode, ip, 1, mode1)
				p2 := getParameter(&intcode, ip, 2, mode2)
				if p1 == 0 {
					ip = p2
					s = 0
				}
			}
		case 7:
			{ // Less than
				s = 4
				p1 := getParameter(&intcode, ip, 1, mode1)
				p2 := getParameter(&intcode, ip, 2, mode2)
				if p1 < p2 {
					setParameter(&intcode, ip, 3, mode3, 1)
				} else {
					setParameter(&intcode, ip, 3, mode3, 0)
				}
			}
		case 8:
			{ // Equals
				s = 4
				p1 := getParameter(&intcode, ip, 1, mode1)
				p2 := getParameter(&intcode, ip, 2, mode2)
				if p1 == p2 {
					setParameter(&intcode, ip, 3, mode3, 1)
				} else {
					setParameter(&intcode, ip, 3, mode3, 0)
				}
			}
		case 99:
			{
				s = 0
				run = false
			}
		}

		ip += s
	}

	return intcode
}

func splitOpcode(n int) (int, int, int, int) {
	padded := fmt.Sprintf("%d", n)
	for len(padded) < 5 {
		padded = "0" + padded
	}

	// fmt.Printf("Padded: %s\n", padded)

	op, _ := strconv.ParseInt(padded[3:5], 10, 32)
	mode1, _ := strconv.ParseInt(padded[2:3], 10, 32)
	mode2, _ := strconv.ParseInt(padded[1:2], 10, 32)
	mode3, _ := strconv.ParseInt(padded[0:1], 10, 32)

	// fmt.Printf("Op: %d\n", op)
	// fmt.Printf("Mode 1: %d\n", mode1)
	// fmt.Printf("Mode 2: %d\n", mode2)
	// fmt.Printf("Mode 3: %d\n", mode3)

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
