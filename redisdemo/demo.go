package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m1 = sync.Mutex{}
	m2 = sync.Mutex{}
)

func goroutine1() {
	m1.Lock()
	m2.Lock()

	time.Sleep(1 * time.Second)
	fmt.Println("11111111")

	m1.Unlock()
	m2.Unlock()
}

func goroutine2() {
	m2.Lock()
	m1.Lock()

	time.Sleep(2 * time.Second)
	fmt.Println("2222222")

	m2.Unlock()
	m1.Unlock()
}

func main() {
	//c, err := redis.Dial("tcp", ":6379")
	//if err != nil {
	//	fmt.Println("link error")
	//	return
	//}
	//c.Do("RPush", "famous", "红楼梦", "三国演义", "西游记", "水浒传")
	//reply, err := redis.Strings(c.Do("LRange", "famous", 0, -1))
	//if err != nil {
	//	fmt.Println("LRange error")
	//	return
	//}
	//fmt.Println(reply)

	for i := 0; i < 3; i++ {
		go goroutine1()
		go goroutine2()
	}

	time.Sleep(60 * time.Second)
}
