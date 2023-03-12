package diffsquares

func SquareOfSum(n int) int {
	return (n * (1 + n) / 2) * (n * (1 + n) / 2)
	// (n * (1 + n) / 2).pow(2)
}

func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
