package file

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// 242ms

var (
	query          = "package.json"
	matched        = 0
	workerCount    = 0
	maxWorkerCount = 32
	searchRequest  = make(chan string)
	workDone       = make(chan bool)
	foundMatch     = make(chan bool)
)

func TestFind(t *testing.T) {
	start := time.Now()
	workerCount = 1
	go search("/Users/haven/code/lowcode-editor/", true)
	waitWorker()
	fmt.Println("match", matched)
	fmt.Println(time.Since(start))
}

func waitWorker() {
	for {
		select {
		case path := <-searchRequest:
			workerCount++
			go search(path, true)
		case <-workDone:
			workerCount--
			if workerCount == 0 {
				return
			}
		case <-foundMatch:
			matched++
		}
	}
}

func search(path string, master bool) {
	files, err := os.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == query {
				foundMatch <- true
			}
			if file.IsDir() {
				dirPath := path + name + "/"
				if workerCount > maxWorkerCount {
					search(dirPath, false)
				} else {
					searchRequest <- dirPath
				}
			}
		}
		if master {
			workDone <- true
		}
	}
}
