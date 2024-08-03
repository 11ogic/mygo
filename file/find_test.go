package file

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var (
	matches = 0
	query   = "package.json"
)

func TestFind(t *testing.T) {
	start := time.Now()
	searchFile("/Users/haven/code/keep_and_increase")
	fmt.Printf("matches: %d \n", matches)
	fmt.Println(start)
}

func searchFile(path string) {
	files, err := os.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == query {
				matches++
			}
			if file.IsDir() {
				searchFile(path + "/" + name)
			}
		}
	}
}
