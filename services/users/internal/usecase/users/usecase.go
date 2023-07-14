package users

import (
	"cinematic.back/services/users/internal/domain/models/user"
	"cinematic.back/services/users/internal/domain/services"
	"context"
	"github.com/google/uuid"
)

type UseCase struct {
	uServices services.UserService
}

func (u UseCase) FindUsersByIds(ctx context.Context, ids ...uuid.UUID) ([]*user.User, error) {
	return u.uServices.FindUsersByIds(ctx, ids...)
}

func (u UseCase) Create(ctx context.Context, data *user.User) (*user.User, error) {
	if err := u.uServices.Save(ctx, data); err != nil {
		return nil, err
	}

	return data, nil
}

func (u UseCase) UpdateByID(ctx context.Context, id uuid.UUID, data *user.User) (*user.User, error) {
	data.ID = id
	if err := u.uServices.Save(ctx, data); err != nil {
		return nil, err
	}

	return data, nil
}

func (u UseCase) DeleteUserByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	return u.uServices.DeleteByID(ctx, id)
}

func (u UseCase) GetUserByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	return u.uServices.FindByID(ctx, id)
}

func NewUseCase(uServices services.UserService) *UseCase {
	return &UseCase{uServices: uServices}
}
