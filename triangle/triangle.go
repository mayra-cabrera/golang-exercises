package triangle

import (
	"math"
	"sort"
)

const testVersion = 2

// Code this function.
func KindFromSides(a, b, c float64) Kind {
	sides := [3]float64{a, b, c}
	unique_sides := removeDuplicated(sides)
	var finalType Kind
	if (triangleEquality(sides)) && unique_sides[0] > 0.0 && !hasInfinity(sides) {
		if len(unique_sides) == 1 {
			finalType = Equ
		}
		if len(unique_sides) == 2 {
			finalType = Iso
		}
		if len(unique_sides) == 3 {
			finalType = Sca
		}
	}
	return finalType
}

func triangleEquality(sides [3]float64) bool {
	return (sides[0]+sides[1] >= sides[2]) && (sides[1]+sides[2] >= sides[0]) && (sides[2]+sides[0] >= sides[1])
}

func hasInfinity(sides [3]float64) bool {
	result := false
	for i := 0; i < len(sides); i++ {
		if math.IsInf(sides[i], +1) || math.IsInf(sides[i], -1) {
			result = true
		}
	}
	return result
}

// Notice it returns this type.  Pick something suitable.
type Kind int

// Pick values for the following identifiers used by the test program.
const (
	NaT = 0 // not a triangle
	Equ = 1 // equilateral
	Iso = 2 // isosceles
	Sca = 3 // scalene
)

// Removes duplicated elements from a float64 slice and sorts them
func removeDuplicated(elements [3]float64) []float64 {
	encountered := map[float64]bool{}

	for v := range elements {
		encountered[elements[v]] = true
	}

	result := []float64{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	sort.Float64s(result)
	return result
}
