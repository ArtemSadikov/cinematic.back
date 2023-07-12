package resolvers

import (
	"cinematic.back/api/users"
	"cinematic.back/services/gateway/internal/services/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	uClient  users.Client
	uService user.Service
}

func NewResolver(
	uClient users.Client,
	uService user.Service,
) *Resolver {
	return &Resolver{uClient, uService}
}
