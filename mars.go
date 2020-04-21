package mars

var (
	DefaultVersion = "0.01"
)

type (
	Mars interface {
		Init(...Option) error
		Run() error
		//Start() error
		Stop() error
		String() string
		Name() string
		// 关闭连接
		CloseConnect(connId uint32)
	}

	Option func(*Options)
)

func NewServer(opts ...Option) Mars {
	return newServer(opts...)
}
