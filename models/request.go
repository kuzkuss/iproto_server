package models

type Request struct {
	Header Header
	Body []byte
}

type RequestSaveString struct {
	Idx int
	Str string
}

type RequestGetString struct {
	Idx int
}

