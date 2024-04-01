package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"net"
)

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

func sendMessage(message string, ip string) {
	fmt.Printf("message: %v ---%v \n", message, ip)
}

func registry(conn net.Conn) bool {
	var (
		result = make([]byte, 1024)
	)

	for i := 0; i < 2; i++ {
		if i == 0 {
			conn.Write([]byte("please enter your username"))
		} else {
			conn.Write([]byte("please enter your password"))
		}
		n, err := conn.Read(result)
		if err != nil {
			fmt.Println("registry error")
			break
		}
		message := string(result[:n])
		username := ""

		for {
			redisConn := pool.Get()
			if i == 0 {
				_, err := redisConn.Do("HSet", message, "username", message)
				if err != nil {
					conn.Write([]byte("username is wrong"))
					fmt.Println(err)
					return false
				}
				username = message
			} else {
				_, err := redisConn.Do("HSet", username, "password", message)
				if err != nil {
					conn.Write([]byte("password is wrong"))
					conn.Close()
					return false
				}
				conn.Write([]byte("welcome to my mind palace \n"))
				return true
			}
		}
	}
	return false
}

func login(conn net.Conn) bool {
	var (
		result = make([]byte, 1024)
	)
	redisConn := pool.Get()
	defer redisConn.Close()

	for i := 0; i < 2; i++ {
		if i == 0 {
			conn.Write([]byte("please enter your username"))
		} else {
			conn.Write([]byte("please enter your password"))
		}

		n, err := conn.Read(result)
		if err != nil {
			fmt.Println("login error")
			break
		}
		message := string(result[:n])
		username := ""

		for {
			if i == 0 {
				reply, err := redis.String(redisConn.Do("HGet", message, "username"))
				if err != nil {
					conn.Write([]byte("username is wrong"))
					return false
				}
				username = reply
			} else {
				password, err := redis.String(redisConn.Do("HGet", username, "password"))
				if err != nil {
					conn.Write([]byte("password is wrong"))
					conn.Close()
					return false
				}
				if password == message {
					conn.Write([]byte("welcome to my mind palace \n"))
					return true
				}
			}
		}
	}
	return false
}

func process(conn net.Conn) {
	for {
		state := "log out"
		result := make([]byte, 1024)
		n, err := conn.Read(result)
		if err != nil {
			fmt.Println("tcp close")
			return
		}
		message := string(result[:n])
		ip := conn.RemoteAddr().String()

		if state == "log out" {
			if message == "login" {
				success := login(conn)
				if success {
					state = "log in"
				}
			} else {
				registry(conn)
			}
		} else {
			sendMessage(message, ip)
		}
	}
	defer conn.Close()
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
		go process(conn)
	}
}
