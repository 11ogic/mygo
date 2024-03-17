package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestIo(t *testing.T) {
	channel := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 2)
		channel <- 5
	}()
	fmt.Println("test: ", <-channel)
}
