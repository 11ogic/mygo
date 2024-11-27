package znet

import (
	"fmt"
	"mygo/zinx/zface"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    zface.IRouter
}

//func Echo(conn *net.TCPConn, buf []byte, size int) error {
//	if _, err := conn.Write(buf[:size]); err != nil {
//		fmt.Println("write buf back error", err)
//		return err
//	}
//	return nil
//}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP %s, Port %d is starting...\n", s.IP, s.Port)
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve TCP addr error", err)
		return
	}
	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("Listen TCP error", err)
		return
	}
	fmt.Println("start server success")
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("connect failed")
			continue
		}
		var cid uint32
		cid++
		dealConn := NewConnection(conn, cid, s.Router)

		go dealConn.Start()
	}

}

func (s *Server) Stop() {}

func (s *Server) Serve() {
	go s.Start()
	select {}
}

func (s *Server) AddRouter(router zface.IRouter) {
	s.Router = router
}

func NewSever(name string) zface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8099,
		Router:    nil,
	}
}
