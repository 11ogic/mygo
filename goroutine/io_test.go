package goroutine

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
)

var (
	done       = make(chan bool)
	allDone    = make(chan bool)
	sourcePath = "source.txt"
	targetPath = "target.txt"
)

func writeDataToFile() {
	file, err := os.OpenFile(sourcePath, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("file error!")
		return
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < 1000; i++ {
		data := rand.Intn(100)
		_, err := writer.WriteString(fmt.Sprintf("%v\n", data))
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}
	defer func() {
		done <- true
		writer.Flush()
	}()
}

func sort() {
	<-done
	file, err := os.Open(sourcePath)
	if err != nil {
		fmt.Println("file error!")
		return
	}
	content, _ := io.ReadAll(file)
	contentArr := strings.Split(string(content), "\n")
	convIntArr := make([]int, 0)
	for _, v := range contentArr[:len(contentArr)-1] {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		convIntArr = append(convIntArr, num)
	}
	newFile, _ := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE, 0666)
	writer := bufio.NewWriter(newFile)
	for _, v := range quickSort(convIntArr) {
		writer.WriteString(fmt.Sprintf("%v\n", v))
	}
	writer.Flush()
	allDone <- true
}

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	var (
		mid   = len(nums) / 2
		pivot = nums[mid]
		left  = make([]int, 0)
		right = make([]int, 0)
	)
	for idx, val := range nums {
		if idx == mid {
			continue
		}
		if val > pivot {
			right = append(right, val)
		} else {
			left = append(left, val)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func TestIo(t *testing.T) {
	go writeDataToFile()
	go sort()
	<-allDone
}
