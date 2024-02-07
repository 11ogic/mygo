package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(filename string) string {
	return func(filename string) string {
		if strings.HasSuffix(filename, suffix) {
			return filename
		} else {
			return filename + suffix
		}
	}
}

func main() {
	handleFilename := makeSuffix(".jpg")

	fmt.Println(handleFilename("caret.jp"))
}
