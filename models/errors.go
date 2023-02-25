package models

import "errors"

var (
	ErrIndex = errors.New("index out of range")
	ErrFuncId = errors.New("invalid request func")
	ErrState = errors.New("invalid state")
	ErrAccess = errors.New("access denied")
)

var errorCodes = map[error]int{
	ErrIndex: 1,
	ErrFuncId: 2,
	ErrState: 3,
	ErrAccess: 4,
}