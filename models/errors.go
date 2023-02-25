package models

import "errors"

var (
	ErrIndex = errors.New("index out of range")
	ErrFuncId = errors.New("invalid request func")
	ErrState = errors.New("invalid state")
)

var errorCodes = map[int]error{
	0x00000001: ErrIndex,
	0x00000002: ErrFuncId,
	0x00000003: ErrState,
}