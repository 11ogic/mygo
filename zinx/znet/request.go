package znet

import "net"

type Request struct {
	Conn *net.TCPConn
	Data []byte
}

func (r *Request) GetData() []byte {
	return r.Data
}

func (r *Request) GetConn() net.TCPConn {
	return *r.Conn
}
