package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var (
	numMap = make(map[int]uint64, 10)
	mutex  sync.Mutex
)

func factorialTailRec(n int, acc uint64) uint64 {
	if n == 0 {
		return acc
	}
	return factorialTailRec(n-1, acc*uint64(n))
}

func factorial(n int) {
	result := factorialTailRec(n, 1)
	mutex.Lock()
	numMap[n] = result
	mutex.Unlock()
}

func TestFactorial(t *testing.T) {
	for i := 1; i <= 65; i++ {
		go factorial(i)
	}
	for k, v := range numMap {
		mutex.Lock()
		fmt.Println(k, "->", v)
		mutex.Unlock()
	}
}
