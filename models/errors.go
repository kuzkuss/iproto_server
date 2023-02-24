package models

import "errors"

var (
	ErrIndex = errors.New("index out of range")
	// ReadError  = errors.New("can not read data from connection")
	// WriteError = errors.New("can not write to connection")
	// DialError  = errors.New("can not connect to address")
	// CloseError = errors.New("error close connection")
)