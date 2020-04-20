package server

import (
	"fmt"
	"net"
)

type (
	tcpServer struct {
		opts     Options
		listener net.Listener
	}
)

func (t *tcpServer) Init(opts ...Option) {
	options := newOptions()
	for _, o := range opts {
		o(&options)
	}
}

func (t *tcpServer) Run() error {
	listener, err := net.Listen(t.opts.Net, fmt.Sprint("%s:%d", t.opts.Addr, t.opts.Port))
	if err != nil {
		return err
	}
	go func() {
		for {
			_, err := listener.Accept()
			if err != nil {
				//TODO 通知主进程退出
				_ = listener.Close()
			}
			// TODO 通知主进程有新的客户端进入
		}
	}()
	return nil
}

func (t *tcpServer) Stop() error {
	panic("implement me")
}

func (t *tcpServer) String() string {
	return "tcp server"
}

func NewTcpServer() Server {
	return &tcpServer{}
}

func newOptions() Options {
	return Options{
		Addr: "",
		Port: 0,
		Net:  "tcp",
	}
}
