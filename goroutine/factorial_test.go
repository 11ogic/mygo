package goroutine

import (
	"fmt"
	"testing"
)

var (
	numMap = make(map[int]int, 200)
)

func factorialTailRec(n int, acc int) int {
	if n == 0 {
		return acc
	}
	return factorialTailRec(n-1, acc*n)
}

func factorial(n int) int {
	return factorialTailRec(n, 1)
}

func TestFactorial(t *testing.T) {
	for i := 1; i <= 200; i++ {
		result := factorial(i)
		fmt.Println(result)
	}
}
