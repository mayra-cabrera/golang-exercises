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
	for i := 2; i <= x; i++ {
		counting := 0
		for j := 2; j <= i; j++ {
			if (i % j) == 0 {
				counting += 1
				if (i == j) && counting == 1 {
					primes = append(primes, i)
				}
			}
		}

	}
	return primes
}

// The test program has a benchmark too.  How fast does your Convert convert?
