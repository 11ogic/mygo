package file

import (
	"fmt"
	"testing"
	"time"
)

var (
	matched = 0
)

func TestFind(t *testing.T) {
	start := time.Now()
	// TODO: Search
	fmt.Println("matched: ", matched)
	fmt.Println(time.Since(start))
}
