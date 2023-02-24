package usecase

import (
	"github.com/kuzkuss/iproto_server/internal/repository"
)

type UseCaseI interface {
	
}

type useCase struct {
	rep repository.RepositoryI
}

func New(rep repository.RepositoryI) UseCaseI {
	return &useCase{
		rep: rep,
	}
}

