package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	wg   = sync.WaitGroup{}
	exit = make(chan struct{})
)

func TestContext(t *testing.T) {
	//ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			ip, err := getIP(exit)
			fmt.Println(ip, err, "@@@")
		}()
	}
	go func() {
		time.Sleep(2 * time.Second)
		close(exit)
		//exit <- struct{}{}
	}()
	wg.Wait()
	fmt.Println("Done...")
}

func getIP(exit chan struct{}) (ip string, err error) {
	go func() {
		select {
		case <-exit:
			wg.Done()
			fmt.Println("Exit...")
			//err = ctx.Err()
			return
		}
	}()
	time.Sleep(4 * time.Second)
	ip = "0.0.0.0"
	fmt.Println("Get IP...")
	wg.Done()
	return ip, nil
}
