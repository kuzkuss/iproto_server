package usecase

import (
	"github.com/kuzkuss/iproto_server/internal/repository"
	"github.com/kuzkuss/iproto_server/models"
)

type UseCaseI interface {
	SwitchState(state int) error
	SaveString(idx int, str string) error
	GetString(idx int) (string, error)
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

func (uc *useCase) SaveString(idx int, str string) error {
	if uc.rep.GetState() != models.READ_WRITE {
		return models.ErrAccess
	}

	err := uc.rep.SaveString(idx, str)
	return err
}

func (uc *useCase) GetString(idx int) (string, error) {
	if uc.rep.GetState() == models.MAINTENANCE {
		return "", models.ErrAccess
	}

	str, err := uc.rep.GetString(idx)
	return str, err
}

