package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"reflect"
)

func ReadData(c net.Conn, inter interface{}) (err error) {
	var (
		result = make([]byte, 1024)
		rVal   = reflect.ValueOf(inter)
	)
	n, err := c.Read(result)
	if err != nil {
		panic("An error occurred while reading")
	}
	err = json.Unmarshal(result[:n], rVal)
	if err != nil {
		fmt.Println(string(result[:n]))
		return err
	}
	return
}

func WriteData(c net.Conn, inter interface{}) (err error) {
	return
}
