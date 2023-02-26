package models

import "errors"

var (
	ErrIndex = errors.New("index out of range")
	ErrFuncId = errors.New("invalid request func")
	ErrState = errors.New("invalid state")
	ErrAccess = errors.New("access denied")
	ErrRead  = errors.New("can not read data from connection")
	ErrWrite  = errors.New("can not write data to connection")
	ErrUnmarshal  = errors.New("can not unmarshal data from request")
)

var ErrorCodes = map[error]uint32{
	ErrIndex: 1,
	ErrFuncId: 2,
	ErrState: 3,
	ErrAccess: 4,
	ErrRead: 5,
	ErrWrite: 6,
	ErrUnmarshal: 7,
}