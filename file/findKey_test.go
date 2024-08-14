package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

const (
	targetDir = "/Users/monterey/code/locales-web/"
)

func contains(slice []string, query string) bool {
	for _, item := range slice {
		if item == query {
			return true
		}
	}
	return false
}
func getTargets(path string) []string {
	var result []string
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return result
	}
	for _, file := range files {
		if !file.IsDir() {
			result = append(result, file.Name())
		}
	}
	return result
}
func hasDuplicateKey(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("[ERR] path is not exist")
		return ""
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var (
		keys    []string
		inBlock = false
	)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading line:", err)
			}
			break
		}
		if inBlock {
			if strings.HasPrefix(strings.TrimSpace(string(line)), "}") {
				inBlock = false
			} else {
				continue
			}
		}
		if strings.Contains(string(line), ":") {
			parts := strings.SplitN(string(line), ":", 2)
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if contains(keys, key) {
				return key
			} else {
				//fmt.Printf("[DEBUG] added key: %v \n", key)
				keys = append(keys, key)
			}
			if strings.HasPrefix(value, "{") {
				//fmt.Printf("[DEBUG]Skipping block %v", string(line))
				inBlock = true
			}
		}
	}
	return ""
}
func TestFindKey(t *testing.T) {
	start := time.Now()
	targets := getTargets(targetDir)
	for _, path := range targets {
		duplicateKey := hasDuplicateKey(targetDir + path)
		if len(duplicateKey) != 0 {
			fmt.Printf("\n[ERR]Duplicate key: '%v' in '%v' \n\n", duplicateKey, path)
			break
		} else {
			fmt.Printf("Finished: %v \n", path)
		}
	}
	fmt.Printf("time: %v", time.Since(start))
}
