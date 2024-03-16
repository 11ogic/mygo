package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var (
	numChan = make(chan int, 2000)
	resChan = make(chan int, 2000)
	wg      sync.WaitGroup
)

func writeNumChan() {
	for i := 1; i <= 2000; i++ {
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
	wg.Done()
}

func TestChannel(t *testing.T) {
	go writeNumChan()
	wg.Add(8)
	for i := 0; i < 8; i++ {
		go writeResChan()
	}
	wg.Wait()
	close(resChan)

	i := 0
	for v := range resChan {
		v, ok := <-resChan
		if !ok {
			break
		}
		fmt.Printf("res[%v] = %v\n", i, v)
		i++
	}
}
