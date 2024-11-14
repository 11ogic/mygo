package znet

import (
	"fmt"
	"io"
	"mygo/zinx/zface"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

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
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connect failed")
			continue
		}
		go func() {
			for {
				buf := make([]byte, 512)
				cnt, err := conn.Read(buf)
				if err != nil {
					fmt.Println("conn read data err", err)
					if err == io.EOF {
						return
					}
					continue
				}
				if _, err := conn.Write(buf[:cnt]); err != nil {
					fmt.Println("write buf back error", err)
					return
				}
			}
		}()
	}

}

func (s *Server) Stop() {}

func (s *Server) Serve() {
	go s.Start()
	select {}
}

func NewSever(name string) zface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8099,
	}
}
