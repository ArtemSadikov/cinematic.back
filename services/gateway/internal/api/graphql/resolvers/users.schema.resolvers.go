package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"errors"

	"cinematic.back/services/gateway/internal/api/graphql/model"
)

// Me is the resolver for the me field.
func (r *mutationResolver) Me(ctx context.Context) (*model.User, error) {
	usr, ok := r.uService.FromIncomingCtx(ctx)
	if !ok {
		return nil, errors.New("no user")
	}

	return &model.User{
		ID: usr.Id.String(),
		Profile: &model.UserProfile{
			Email:    usr.Profile.Email,
			Username: usr.Profile.Username,
		},
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
		DeletedAt: usr.DeletedAt,
	}, nil
}
