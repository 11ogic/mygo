package main

import "mygo/zinx/znet"

func main() {
	server := znet.NewSever("Demo 1")
	server.Serve()
}
