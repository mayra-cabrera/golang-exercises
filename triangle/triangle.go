package triangle

const testVersion = 2

// Code this function.
func KindFromSides(a, b, c float64) Kind {
	sides := [3]float64{a, b, c}
	unique_sides := removeDuplicated(sides)
	var finalType Kind
	if sides[0]+sides[1] >= sides[2] {
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

// Notice it returns this type.  Pick something suitable.
type Kind int

// Pick values for the following identifiers used by the test program.
const (
	NaT = 0 // not a triangle
	Equ = 1 // equilateral
	Iso = 2 // isosceles
	Sca = 3 // scalene
)

func removeDuplicated(elements [3]float64) []float64 {
	encountered := map[float64]bool{}

	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []float64{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

// Organize your code for readability.
