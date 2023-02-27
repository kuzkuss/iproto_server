package repository

import (
	"github.com/kuzkuss/iproto_server/models"
)

type RepositoryI interface {
	ChangeState(newState int)
	SaveString(idx int, str string) error
	GetString(idx int) (string, error)
	GetState() int
}

const SIZE_STORE = 1000

type storage struct {
	strs [SIZE_STORE]string
	state int
}

func New() RepositoryI {
	return &storage {
		state: models.READ_WRITE,
	}
}

func (strg *storage) ChangeState(newState int) {
	strg.state = newState
}

func (strg *storage) GetState() int {
	return strg.state
}

func (strg *storage) SaveString(idx int, str string) error {
	if idx >= SIZE_STORE || idx < 0 {
		return models.ErrIndex
	}
	strg.strs[idx] = str
	return nil
}

func (strg *storage) GetString(idx int) (string, error) {
	if idx >= SIZE_STORE || idx < 0 {
		return "", models.ErrIndex
	}
	return strg.strs[idx], nil
}
