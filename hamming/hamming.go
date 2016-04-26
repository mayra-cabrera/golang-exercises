package hamming

import "errors"

const testVersion = 4

func Distance(a, b string) (int, error) {
	hammingValue := 0
	dna1 := []rune(a)
	dna2 := []rune(b)
	if len(dna1) != len(dna2) {
		return -1, errors.New("Sequences have differente lengths!")
	}
	for i, v := range dna1 {
		if v != dna2[i] {
			hammingValue += 1
		}
	}
	return hammingValue, nil
}
