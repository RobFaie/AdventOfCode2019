package lib

// RunIntcode processes an intcode
func RunIntcode(intcode []int64) []int64 {

	var (
		ip  int64
		s   int64
		run bool = true
	)

	for run {
		switch op := intcode[ip]; op {
		case 1:
			{ // Addition
				s = 4
				i1 := intcode[intcode[ip+1]]
				i2 := intcode[intcode[ip+2]]
				intcode[intcode[ip+3]] = i1 + i2
			}
		case 2:
			{ // Multiplication
				s = 4
				i1 := intcode[intcode[ip+1]]
				i2 := intcode[intcode[ip+2]]
				intcode[intcode[ip+3]] = i1 * i2
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
