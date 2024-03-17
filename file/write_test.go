package file

import (
	"fmt"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	file, err := os.OpenFile("./test.txt", os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	if err == nil {
		file.WriteString("Hello World")
	} else {
		fmt.Println("err is ", err)
	}
}
