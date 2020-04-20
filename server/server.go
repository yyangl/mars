package server

type (
	Server interface {
		Init(...Option)
		Run() error
		Stop() error
		String() string
	}

	Conn interface {
		Send([]byte) (int, error)
		//Read() ([]byte, error)
		GetID() uint32
		Run() error
	}
	HandlerFunc func(Conn)

	Option func(*Options)
)

var (
	DefaultServer = NewTcpServer()
)
