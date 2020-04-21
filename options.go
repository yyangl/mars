package mars

import (
	"context"
	"github.com/yyangl/mars/codec"
	"github.com/yyangl/mars/server"
)

type Options struct {
	Version string
	Server  server.Server
	Codec   codec.Codec
	Name    string
	Addr    string
	Port    int
	Signal  bool

	Before []func() error
	After  []func() error

	Context context.Context
}

func Addr(addr string) Option {
	return func(o *Options) {
		o.Addr = addr
	}
}
