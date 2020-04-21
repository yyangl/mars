package mars

import (
	"context"
	"fmt"
	"github.com/yyangl/mars/codec"
	"github.com/yyangl/mars/connect"
	"github.com/yyangl/mars/server"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type (
	marsServer struct {
		opts     Options
		once     sync.Once
		connects map[uint32]connect.Connect
		index    uint32
		mux      sync.Mutex
		// conn chan 接受连接的客户端
		cc chan connect.Connect
	}
)

func (m *marsServer) HandlerConn() {
	go func() {
		for {
			conn := <-m.cc
			m.mux.Lock()
			m.connects[m.index] = conn
			m.index++
			m.mux.Unlock()
		}
	}()
}

func (m *marsServer) Codec() codec.Codec {
	panic("implement me")
}

func (m *marsServer) CloseConnect(connId uint32) {
	m.mux.Lock()
	defer m.mux.Unlock()
	delete(m.connects, connId)
}

func (m *marsServer) Init(opts ...Option) error {
	for _, o := range opts {
		o(&m.opts)
	}
	m.once.Do(func() {

	})
	return nil
}

func (m *marsServer) Name() string {
	return m.opts.Name
}

func (m *marsServer) Run() error {
	for _, fn := range m.opts.Before {
		if err := fn(); err != nil {
			return err
		}
	}

	if err := m.opts.Server.Run(m.opts.Context); err != nil {
		return err
	}
	ch := make(chan os.Signal, 1)
	if m.opts.Signal {
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	}

	//
	m.HandlerConn()

	select {
	case <-ch:
		//case <-m.opts.Context.Do():
	case <-m.opts.Context.Done():
		fmt.Printf("[Mars] tcp server exit...")
	}
	return m.Stop()
}

func (m *marsServer) Stop() error {
	//close(m.cc)
	return m.opts.Server.Stop()
}

func (m *marsServer) String() string {
	return "mars"
}

func newOptions(opts ...Option) Options {
	options := Options{
		Version: DefaultVersion,
		Server:  server.DefaultServer,
		Context: context.Background(),
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

func newServer(opts ...Option) Mars {
	m := &marsServer{
		cc:       make(chan connect.Connect, 10),
		connects: make(map[uint32]connect.Connect),
	}
	options := newOptions(opts...)
	options.Server.Init(
		server.Addr(options.Addr),
		server.Port(options.Port),
		server.CodecHandler(options.Codec),
		server.CChanHandler(m.cc),
	)
	m.opts = options
	return m
}
