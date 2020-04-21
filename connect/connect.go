package connect

type Connect interface {
	Send([]byte) (int, error)
	GetID() uint32
	Run() error
}
