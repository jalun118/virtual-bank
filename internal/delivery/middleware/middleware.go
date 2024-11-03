package middleware

import "virtual-bank/internal/usecase"

type Middlewares struct {
	sessionUsecase *usecase.SessionUsecase
	userUsecase    *usecase.UserUsecase
}

func NewMiddlewares(sessionUsecase *usecase.SessionUsecase, userUsecase *usecase.UserUsecase) *Middlewares {
	return &Middlewares{
		sessionUsecase: sessionUsecase,
		userUsecase:    userUsecase,
	}
}
