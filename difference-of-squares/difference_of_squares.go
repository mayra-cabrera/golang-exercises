package diffsquares

func SquareOfSums(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		result += i
	}
	return result * result
}

func SumOfSquares(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		result += i * i
	}
	return result
}

func Difference(n int) int {
	squareOfSum := SquareOfSums(n)
	sumOfSquare := SumOfSquares(n)
	return squareOfSum - sumOfSquare
}
