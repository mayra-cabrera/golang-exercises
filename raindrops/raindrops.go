package raindrops

import "strconv"

const testVersion = 2

func Convert(x int) string {
	result := ""
	primes := SearchPrimeNumbers(x)
	for _, v := range primes[:] {
		word := ""
		switch v {
		case 3:
			word = "Pling"
		case 5:
			word = "Plang"
		case 7:
			word = "Plong"
		}
		result = result + word
	}

	if len(result) == 0 {
		result = strconv.Itoa(x)
	}
	return result
}

func SearchPrimeNumbers(x int) []int {
	primes := []int{}
	p := 2
	if x < 2 {
		primes = append(primes, p)
	} else {
		for p <= x {
			if x%p == 0 {
				primes = append(primes, p)

			}
			p += 1
		}
	}

	return primes
}

// The test program has a benchmark too.  How fast does your Convert convert?
