package models

type Response struct {
	Header Header
	ReturnCode int
	Body []byte
}

