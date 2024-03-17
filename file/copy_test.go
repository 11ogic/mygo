package file

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func Copy(p1 string, p2 string) (int64, error) {
	reader, _ := os.Open(p1)
	fmt.Println(reader)
	defer reader.Close()
	writer, _ := os.OpenFile(p2, os.O_WRONLY|os.O_CREATE, 0666)
	defer writer.Close()
	return io.Copy(writer, reader)
}

func TestCopy(t *testing.T) {
	Copy("./test.txt", "./text2.txt")
}
