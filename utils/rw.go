package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"reflect"
)

func ReadData(c net.Conn, inter interface{}) (err error) {
	if reflect.TypeOf(inter).Kind() != reflect.Ptr {
		return errors.New("inter must be a pointer")
	}
	result := make([]byte, 1024)
	n, err := c.Read(result)
	if err != nil {
		panic("An error occurred while reading")
	}
	err = json.Unmarshal(result[:n], inter)
	if err != nil {
		fmt.Println(string(result[:n]))
		return err
	}
	return
}

func WriteData(c net.Conn, inter interface{}) (err error) {
	return
}
