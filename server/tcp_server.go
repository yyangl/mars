package server

import (
	"context"
	"fmt"
	"github.com/yyangl/mars/connect"
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
	t.opts = options
}

func (t *tcpServer) Run(ctx context.Context) error {
	fmt.Println("[Mars] server run ...")
	listener, err := net.Listen(t.opts.Net, fmt.Sprintf("%s:%d", t.opts.Addr, t.opts.Port))
	if err != nil {
		return err
	}
	fmt.Printf("[Mars] server run addr %s", listener.Addr().String())
	go func() {
		for {
			conn, err := listener.Accept()
			if err == nil {
				//通知主进程退出
				//_ = listener.Close()
				fmt.Printf("[Mars] accept client err %v\n", err)
				ctx.Done()
				break
			}
			// 通知主进程有新的客户端进入
			t.opts.CChan <- connect.NewTcpConn(conn.(net.Conn))
		}
	}()
	return nil
}

func (t *tcpServer) Stop() error {
	return t.listener.Close()
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
		Net:  "tcp4",
	}
}
