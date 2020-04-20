package server

type Options struct {
	Addr string
	Port int
	Net  string
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
