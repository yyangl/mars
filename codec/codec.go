package codec

type Codec interface {
	GetDataLen() uint
	GetLen() uint
	GetData() []byte
	SetLen() uint
	SetData() uint
}
