package models

type Header struct {
	FuncId uint32
	BodyLength uint32
	RequestId uint32
}

type Request struct {
	Header Header
	Body []byte
}

