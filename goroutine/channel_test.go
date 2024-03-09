package goroutine

import (
	"fmt"
	"testing"
)

var (
	numChan = make(chan int, 2000)
	resChan = make(chan int, 2000)
)

func writeNumChan() {
	for i := 1; i <= cap(numChan); i++ {
		numChan <- i
	}
}

func writeResChan() {
	for v := range numChan {
		sum := 0
		for i := 1; i <= v; i++ {
			sum += i
		}
		resChan <- sum
	}
}

func TestChannel(t *testing.T) {
	go writeNumChan()
	for i := 0; i < 8; i++ {
		go writeResChan()
	}
	i := 0
	for v := range resChan {
		fmt.Printf("res[%v] = %v\n", i, v)
		i++
	}
}
