package server

import (
	"github.com/yyangl/mars/codec"
	"github.com/yyangl/mars/connect"
)

type Options struct {
	Addr   string
	Port   int
	Net    string
	Codec  codec.Codec
	CChan  chan connect.Connect
	Signal chan bool
}

func Port(port int) Option {
	return func(o *Options) {
		o.Port = port
	}
}

func Addr(addr string) Option {
	return func(o *Options) {
		o.Addr = addr
	}
}

func Net(net string) Option {
	return func(o *Options) {
		o.Net = net
	}
}
func CodecHandler(codec codec.Codec) Option {
	return func(o *Options) {
		o.Codec = codec
	}
}

func CChanHandler(c chan connect.Connect) Option {
	return func(o *Options) {
		o.CChan = c
	}
}
