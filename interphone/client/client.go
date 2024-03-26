package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	message := make([]byte, 1024)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte(line))

		n, err := conn.Read(message)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(message[:n]))
	}
}
