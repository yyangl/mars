package connect

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type tcpConn struct {
	id   uint32
	conn net.Conn
	mux  sync.Mutex
	//buf    []byte
}

func NewTcpConn(conn net.Conn) *tcpConn {
	return &tcpConn{
		conn: conn,
		mux:  sync.Mutex{},
	}
}

func (t *tcpConn) Send(bytes []byte) (int, error) {
	t.mux.Lock()
	defer t.mux.Unlock()
	n, err := t.conn.Write(bytes)
	return n, err
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
			_, err := t.conn.Read(buf)
			if err == io.EOF {
				//t.server.Close(t.id)
				_ = t.conn.Close()
				break
			} else if err != nil && err != io.EOF {
				fmt.Printf("read data err %v", err)
				_ = t.conn.Close()
				break
			} else {
				//if t.server.Recv(buf[:n]) {
				//
				//}
			}
		}
	}(t)
	return nil
}
