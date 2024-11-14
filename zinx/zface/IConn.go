package zface

import "net"

type IConnection interface {
	Start()

	Stop()

	GetTCPConnID() uint32

	RemoteAddr() net.Addr

	Send([]byte) error

	GetConn() *net.TCPConn
}

type HandleFunc func(*net.TCPConn, []byte, int) error
