package lib

import (
	"fmt"
	"strconv"
	"strings"
)

// WireSegment describes a wire segment
type WireSegment struct {
	D string
	L int
}

// Wire describes a wire
type Wire []WireSegment

func (ws WireSegment) String() string {
	return fmt.Sprintf("%s:%d", ws.D, ws.L)
}

type position struct {
	X int
	Y int
}

// ParseWire processessa string into a Wire
func ParseWire(s string) Wire {
	// R8,U5,L5,D3
	parts := strings.Split(string(s), ",")

	w := make(Wire, len(parts))
	for i, v := range parts {
		d := string(v[0])
		l, _ := strconv.ParseInt(string(v[1:]), 10, 32)
		w[i] = WireSegment{D: d, L: int(l)}
	}

	return w
}

func procMove(pos position, m map[string]int, hits *[]position, off int, d *map[position]int) {
	posstr := fmt.Sprintf("%d:%d", pos.X, pos.Y)
	m[posstr] += off
	if m[posstr] == 3 {
		*hits = append(*hits, pos)
	}
	// fmt.Printf("%s: %d\n", posstr, m[posstr])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// GetWireOverlapDistance computes the closest overlap between two wires.
func GetWireOverlapDistance(ws []Wire) (int, int) {
	m := make(map[position]int)
	d := make([]map[position]int, len(ws))
	var hits []position

	for i, w := range ws {
		off := 1 << i
		dist := 0
		// fmt.Printf("Offset %d: %d\n", i, off)
		fmt.Printf("Wire %d: %s\n", i, w)

		var pos position = position{X: 0, Y: 0}

		d[i] = make(map[position]int)

		for _, s := range w {
			switch s.D {
			case "R":
				for k := 0; k < s.L; k++ {
					pos.X++
					m[pos] += off
					if m[pos] == 3 {
						hits = append(hits, pos)
					}

					dist++
					if _, b := d[i][pos]; !b {
						d[i][pos] = dist
					}
				}
			case "L":
				for k := 0; k < s.L; k++ {
					pos.X--
					m[pos] += off
					if m[pos] == 3 {
						hits = append(hits, pos)
					}

					dist++
					if _, b := d[i][pos]; !b {
						d[i][pos] = dist
					}
				}
			case "U":
				for k := 0; k < s.L; k++ {
					pos.Y++
					m[pos] += off
					if m[pos] == 3 {
						hits = append(hits, pos)
					}

					dist++
					if _, b := d[i][pos]; !b {
						d[i][pos] = dist
					}
				}
			case "D":
				for k := 0; k < s.L; k++ {
					pos.Y--
					m[pos] += off
					if m[pos] == 3 {
						hits = append(hits, pos)
					}

					dist++
					if _, b := d[i][pos]; !b {
						d[i][pos] = dist
					}
				}
			}
		}
	}

	fmt.Printf("Hits: %s\n", hits)

	mDist1 := 1<<32 - 1
	for _, v := range hits {
		dist := abs(v.X) + abs(v.Y)
		if dist < mDist1 {
			mDist1 = dist
		}
	}

	mDist2 := 1<<32 - 1
	for _, v := range hits {
		dist := d[0][v] + d[1][v]
		if dist < mDist2 {
			mDist2 = dist
		}
	}

	return mDist1, mDist2
}
