package mars

import "github.com/yyangl/mars/server"

type Options struct {
	Version string
	Server  server.Server
	Name    string
	Addr    string
	Port    int

	Before []func() error
	After  []func() error
}
