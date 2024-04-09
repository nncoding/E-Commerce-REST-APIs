package middlewareusecases

import (
	middlewarerepositories "github.com/nncoding/go-basic/modules/middleware/middlewareRepositories"
)

type IMiddlewareUsecase interface {
}

type middlewareUsecase struct {
	middlewareRepository middlewarerepositories.IMiddlewareRepository
}

func MiddlewaresUsecase(middlewareRepository middlewarerepositories.IMiddlewareRepository) IMiddlewareUsecase {
	return &middlewareUsecase{
		middlewareRepository: middlewareRepository,
	}
}
