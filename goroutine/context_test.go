package goroutine

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	wg = sync.WaitGroup{}
)

func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			ip, err := getIP(ctx)
			fmt.Println(ip, err, "@@@")
		}()
	}
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()
	wg.Wait()
	fmt.Println("Done...")
}

func getIP(ctx context.Context) (ip string, err error) {
	go func() {
		select {
		case <-ctx.Done():
			wg.Done()
			err = ctx.Err()
			return
		}
	}()
	time.Sleep(4 * time.Second)
	ip = "0.0.0.0"
	fmt.Println("Get IP...")
	wg.Done()
	return ip, nil
}
