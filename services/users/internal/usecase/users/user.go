package users

import (
	"cinematic.back/services/users/internal/domain/models/user"
	"context"
	"github.com/google/uuid"
)

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
