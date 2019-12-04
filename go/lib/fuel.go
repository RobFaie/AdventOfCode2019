package lib

// FuelForMass returns the fuel required for a given mass.
func FuelForMass(x int64) int64 {
	return x/3 - 2
}
