package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("link error")
		return
	}
	c.Do("RPush", "famous", "红楼梦", "三国演义", "西游记", "水浒传")
	reply, err := redis.Strings(c.Do("LRange", "famous", 0, -1))
	if err != nil {
		fmt.Println("LRange error")
		return
	}
	fmt.Println(reply)
}
