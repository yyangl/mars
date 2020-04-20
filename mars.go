package mars

import (
	"github.com/yyangl/mars/codec"
	"github.com/yyangl/mars/server"
	"sync"
)

var (
	DefaultVersion = "0.01"
)

type (
	Mars interface {
		Init(...Option) error
		Run() error
		Start() error
		Stop() error
		String() string
		Name() string
		Server() server.Server
		Codec() codec.Codec
		// 关闭连接
		Close(connId uint32) error
		// 有新的连接进来
		Connected(server.Conn) error
	}

	mars struct {
		opts     Options
		once     sync.Once
		handler  server.HandlerFunc
		codec    codec.Codec
		connects map[uint32]server.Conn
		index    uint32
		mux      sync.Mutex
	}
	Option func(*Options)
)

func (m *mars) Codec() codec.Codec {
	panic("implement me")
}

func (m *mars) Close(connId uint32) error {
	m.mux.Lock()
	defer m.mux.Unlock()
	delete(m.connects, connId)
	return nil
}

func (m *mars) Connected(server.Conn) error {
	panic("implement me")
}

func (m *mars) Start() error {
	for _, fn := range m.opts.Before {
		if err := fn(); err != nil {
			return err
		}
	}

	if err := m.opts.Server.Run(); err != nil {
		return err
	}
	return nil
}

func (m *mars) Server() server.Server {
	return m.opts.Server
}

func (m *mars) Init(opts ...Option) error {
	for _, o := range opts {
		o(&m.opts)
	}
	m.once.Do(func() {

	})
	return nil
}

func (m *mars) Name() string {
	return m.opts.Name
}

func (m *mars) Run() error {
	return m.opts.Server.Run()
}

func (m *mars) Stop() error {
	return m.opts.Server.Stop()
}

func (m *mars) String() string {
	return "mars"
}

func newOptions(opts ...Option) Options {
	options := Options{
		Version: DefaultVersion,
		Server:  server.DefaultServer,
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

func NewServer(opts ...Option) Mars {
	m := new(mars)
	options := newOptions(opts...)
	options.Server.Init(
		server.Addr(options.Addr),
		server.Port(options.Port),
		server.Mars(m),
	)
	m.opts = options
	return m
}
