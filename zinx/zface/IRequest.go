package zface

import "net"

type IRequest interface {
	GetData() []byte
	GetConn() net.TCPConn
}
