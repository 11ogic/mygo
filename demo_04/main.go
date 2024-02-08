package main

import (
	"errors"
	"fmt"
	"time"
)

func test(b bool) (err error, msg string) {
	if b {
		msg = "failed~"
		err = errors.New("error")
	}
	msg = "success!"
	return
}

func main() {
	err, msg := test(true)
	//if err != nil {
	//	panic(msg)
	//}
	fmt.Printf("err: %v %T, msg: %v \n", err, err, msg)

	for {
		time.Sleep(time.Second)
		fmt.Println("other code...")
	}
}
