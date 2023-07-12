package user

import (
	"cinematic.back/api/users/pb"
	"context"
)

type Service struct{}

func NewService() Service {
	return Service{}
}

var dataKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func (s *Service) ToOutgoingCtx(ctx context.Context, data *pb.User) context.Context {
	user := User{}
	user.fromPB(data)

	return context.WithValue(ctx, dataKey, user)
}

func (s *Service) FromIncomingCtx(ctx context.Context) (User, bool) {
	user, ok := ctx.Value(dataKey).(User)
	if !ok {
		return User{}, false
	}
	return user, true
}
