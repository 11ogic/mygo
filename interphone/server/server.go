package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"mygo/utils"
	"net"
)

type data struct {
	Msg string `json:"msg"`
}

var (
	pool *redis.Pool
)

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379")
		},
	}
}

func test(c net.Conn) {
	mes := utils.NewMessage(c)
	for {
		var req data
		err := mes.ReadData(&req)
		if err != nil {
			fmt.Println("error!!", err)
			break
		}
		fmt.Println("req.msg = ", req.Msg)
	}
	defer c.Close()
}

func main() {
	fmt.Println("link start...")
	l, err := net.Listen("tcp", ":8080")
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
		go test(conn)
	}
}
