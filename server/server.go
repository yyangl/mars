package server

import "context"

type (
	Server interface {
		Init(...Option)
		Run(context.Context) error
		Stop() error
		String() string
	}

	Option func(*Options)
)

var (
	DefaultServer = NewTcpServer()
)
