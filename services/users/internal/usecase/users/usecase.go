package users

import (
	"cinematic.back/services/users/internal/domain/services"
)

type UseCase struct {
	uServices services.UserService
}

func NewUseCase(uServices services.UserService) *UseCase {
	return &UseCase{uServices: uServices}
}
