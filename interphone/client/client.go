package main

import "net"

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		return
	}

	defer conn.Close()

	for i := 0; i < 10000; i++ {
		conn.Write([]byte{'h', 'e', 'l', 'l', 'o'})
	}

}
