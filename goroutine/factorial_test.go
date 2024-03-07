package goroutine

import (
	"fmt"
	"testing"
)

var (
	numMap = make(map[int]uint64, 10)
)

func factorialTailRec(n int, acc uint64) uint64 {
	if n == 0 {
		return acc
	}
	return factorialTailRec(n-1, acc*uint64(n))
}

func factorial(n int) {
	result := factorialTailRec(n, 1)
	numMap[n] = result
}

func TestFactorial(t *testing.T) {
	for i := 1; i <= 65; i++ {
		factorial(i)
	}
	fmt.Println(numMap)
}
