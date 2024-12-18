package znet

import (
	"fmt"
	"mygo/zinx/zface"
	"net"
)

type Connection struct {
	ConnID uint32

	Conn *net.TCPConn

	Router zface.IRouter

	IsClosed bool

	ExitChan chan bool
}

func NewConnection(conn *net.TCPConn, id uint32, router zface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   id,
		Router:   router,
		IsClosed: false,
		ExitChan: make(chan bool),
	}
}

func (c *Connection) StartReader() {
	fmt.Println("StartReader is running... remoteAddr is ", c.RemoteAddr())
	defer func() {
		fmt.Printf("Conn ID = %d, Reader is exit, remoteAddr is %d", c.ConnID, c.RemoteAddr())
		c.Stop()
	}()
	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("Read back error, continue...")
			continue
		}
		req := &Request{
			Conn: c.Conn,
			Data: buf[:cnt],
		}

		go func(request zface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(req)
	}
}

func (c *Connection) Start() {
	fmt.Println("start connection... ID is ", c.ConnID)

	go c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Println("Stop connection... ID is ", c.ConnID)
	if c.IsClosed {
		return
	}
	c.IsClosed = true
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	_, err := c.Conn.Write(data)
	return err
}

func (c *Connection) GetConn() *net.TCPConn {
	return c.Conn
}
