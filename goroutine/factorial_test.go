package goroutine

import "testing"

var (
	numMap = make(map[int]int, 200)
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func TestFactorial(t *testing.T) {
	result := factorial(6)
	numMap[6] = result
	t.Logf("result: %v", result)
}
