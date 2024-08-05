package file

import (
	"fmt"
	"os"
	"time"
)

var (
	maxWorkerCount = 16
	workerCount    = 0
	matched        = 0
	query          = "package.json"
	searchRequest  = make(chan string)
	workDone       = make(chan bool)
	foundFile      = make(chan bool)
)

func main() {
	start := time.Now()
	workerCount = 1
	go search("/Users/monterey/code/", true)
	waitForWorker()
	fmt.Printf("matches: %d \n", matched)
	fmt.Printf("time: %v \n", time.Since(start))
}

func waitForWorker() {
	for {
		select {
		case path := <-searchRequest:
			go search(path, true)
			workerCount++
		case <-foundFile:
			matched++
		case <-workDone:
			workerCount--
			if workerCount == 0 {
				return
			}
		}
	}
}

func search(path string, master bool) {
	defer func() {
		if master {
			workDone <- true
		}
	}()
	files, err := os.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == query {
				foundFile <- true
			}
			if file.IsDir() {
				searchPath := path + name + "/"
				if workerCount > maxWorkerCount {
					searchRequest <- searchPath
				} else {
					search(searchPath, false)
				}
			}
		}
	}
}
