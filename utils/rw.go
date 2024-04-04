package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"reflect"
)

type Message struct {
	C net.Conn
}

func NewMessage(c net.Conn) *Message {
	return &Message{C: c}
}

func (m *Message) ReadData(inter interface{}) (err error) {
	if reflect.TypeOf(inter).Kind() != reflect.Ptr {
		return errors.New("inter must be a pointer")
	}
	buffer := make([]byte, 1024)
	n, err := m.C.Read(buffer[:4])
	if n != 4 || err != nil {
		return errors.New("read fail")
	}
	size := binary.BigEndian.Uint32(buffer[:4])
	request := &RequestType{}
	n, err = m.C.Read(buffer[:size])
	if err != nil || size != uint32(n) {
		return errors.New("read fail")
	}

	err = json.Unmarshal(buffer[:size], request)
	if err != nil {
		fmt.Println(string(buffer[:n]))
		return err
	}
	err = json.Unmarshal([]byte(request.Data), inter)
	if err != nil {
		fmt.Println(string(buffer[:n]))
		return err
	}
	fmt.Printf("size = %d; content = %+v", size, request)
	return
}

func (m *Message) WriteData(inter interface{}) (err error) {
	data, err := json.Marshal(inter)
	if err != nil {
		return errors.New("marshal fail")
	}
	request := &RequestType{Data: string(data), Code: 200}
	sendData, err := json.Marshal(request)
	if err != nil {
		return errors.New("marshal fail")
	}
	size := uint32(len(sendData))
	var sizeBuf [4]byte
	binary.BigEndian.PutUint32(sizeBuf[:4], size)
	n, err := m.C.Write(sizeBuf[:])
	if n != 4 || err != nil {
		fmt.Println(n)
		fmt.Println(err)
		return errors.New("failed to write size")
	}
	n, err = m.C.Write(sendData)
	if n != int(size) || err != nil {
		return errors.New("failed to write sendData")
	}

	return
}
