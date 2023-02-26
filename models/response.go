package models

type Response struct {
	Header Header
	ReturnCode uint32
	Body []byte
}

