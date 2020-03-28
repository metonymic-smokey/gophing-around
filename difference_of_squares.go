package diffsquares

import (
	"fmt"
)

func SquareOfSum(n int) int {

	sum := n * (n + 1) / 2
	sq := sum*sum

	return sq
}

func SumOfSquares(n int) int {
	var i,sum int

	for i = 1; i <= n; i++ {
		sum += i*i
	}

	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)

}

}
