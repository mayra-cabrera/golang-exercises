package raindrops

import "strconv"

const testVersion = 2

func Convert(x int) string {
	result := ""
	for i := 1; i <= x; i++ {
		word := ""
		if (float64(x) / float64(i)) == 1 {
			switch i {
			case 3:
				word = "Pling"
			case 5:
				word = "Plang"
			case 7:
				word = "Plong"
			}
			result = result + word
		}
	}

	if len(result) == 0 {
		result = strconv.Itoa(x)
	}
	return result
}

// The test program has a benchmark too.  How fast does your Convert convert?
