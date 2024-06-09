package file

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	mu             sync.Mutex
	matched        int
	query          = "package.json"
	workerCount    = 0
	maxWorkerCount = 32
	searchRequest  = make(chan string)
	workDone       = make(chan bool)
	foundFile      = make(chan bool)
)

func TestFind(t *testing.T) {
	start := time.Now()
	increaseWorkerCount()
	go search("/Users/", true)
	waitForWorker()
	fmt.Println("matched: ", matched)
	fmt.Println(time.Since(start))
}

func increaseWorkerCount() {
	mu.Lock()
	defer mu.Unlock()
	workerCount++
}

func decrementWorkerCount() {
	mu.Lock()
	defer mu.Unlock()
	workerCount--
}

func getWorkerCount() int {
	mu.Lock()
	defer mu.Unlock()
	return workerCount
}

func waitForWorker() {
	for {
		select {
		case path := <-searchRequest:
			increaseWorkerCount()
			go search(path, true)
		case <-workDone:
			decrementWorkerCount()
			if getWorkerCount() == 0 {
				return
			}
		case <-foundFile:
			matched++
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
				if getWorkerCount() < maxWorkerCount {
					searchRequest <- path + name + "/"
				} else {
					search(path+name+"/", false)
				}
			}
		}
	}
}
