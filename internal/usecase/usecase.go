package usecase

import (
	"github.com/kuzkuss/iproto_server/internal/repository"
	"github.com/kuzkuss/iproto_server/models"
)

type UseCaseI interface {
	SwitchState(state int) error
	SaveString() error
}

type useCase struct {
	rep repository.RepositoryI
}

func New(rep repository.RepositoryI) UseCaseI {
	return &useCase{
		rep: rep,
	}
}

func (uc *useCase) SwitchState(state int) error {
	if state != models.READ_ONLY && state != models.READ_WRITE && state != models.MAINTENANCE {
		return models.ErrState
	}
	uc.rep.ChangeState(state)
	return nil
}

