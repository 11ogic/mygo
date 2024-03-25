package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	for {
		result := make([]byte, 1024)
		n, err := conn.Read(result)
		if err != nil {
			fmt.Println("tcp close")
			return
		}
		_, err = conn.Write([]byte("welcome to my mind palace \n"))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("message: %v            ---%v \n", string(result[:n]), conn.RemoteAddr().String())
	}
	defer conn.Close()
}

func main() {
	fmt.Println("link start...")
	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("link error: ", err)
		return
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go process(conn)
	}
}
