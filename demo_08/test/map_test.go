package test

import (
	"fmt"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	m1 := map[int]int{}
	if v, ok := m1[3]; !ok {
		t.Logf("ok := %v", ok)
	} else {
		fmt.Printf("value := %v", v)
	}

	decorateSlowFn := timeSpent(slowFn)
	res := decorateSlowFn(5)
	fmt.Println(res)
}

func timeSpent(inner func(n int) int) func(n int) int {
	return func(n int) int {
		start := time.Now()
		res := inner(n)
		fmt.Println("used time: ", time.Since(start).Seconds())
		return res
	}
}

func slowFn(n int) int {
	time.Sleep(time.Second * 3)
	return n
}
