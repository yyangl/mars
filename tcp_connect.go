package mars

import (
	"fmt"
	server2 "github.com/yyangl/mars/server"
	"io"
	"net"
	"sync"
)

type tcpConn struct {
	id     uint32
	conn   net.TCPConn
	mux    sync.Mutex
	server server2.Server
	buf    []byte
}

func (t *tcpConn) Send(bytes []byte) (int, error) {
	return t.conn.Write(bytes)
}

//func (t *tcpConn) Read() ([]byte, error) {
//	return t.conn.Read()
//}

func (t *tcpConn) GetID() uint32 {
	return t.id
}

func (t *tcpConn) Run() error {
	go func(t *tcpConn) {
		for {
			buf := make([]byte, 1024)
			n, err := t.conn.Read(buf)
			if err == io.EOF {
				t.server.Close(t.id)
				_ = t.conn.Close()
				break
			} else if err != nil && err != io.EOF {
				fmt.Printf("read data err %v", err)
				_ = t.conn.Close()
				break
			} else {
				if t.server.Recv(buf[:n]) {

				}
			}
		}
	}(t)
	return nil
}
