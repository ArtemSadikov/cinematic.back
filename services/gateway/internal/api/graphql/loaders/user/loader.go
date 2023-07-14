package user

import (
	"cinematic.back/api/users"
	"cinematic.back/services/gateway/internal/api/graphql/model"
	"cinematic.back/services/gateway/internal/services/user"
	"context"
	"github.com/graph-gophers/dataloader"
	"time"
)

const wait = 2 * time.Millisecond

type ctxKey string

const loadersKey = ctxKey("user_dataloader")

type Loader struct {
	uLoader *dataloader.Loader
}

func (l *Loader) GetUser(ctx context.Context, id string) (*model.User, error) {
	thunk := l.uLoader.Load(ctx, dataloader.StringKey(id))
	res, err := thunk()
	if err != nil {
		return nil, err
	}
	return res.(*model.User), nil
}

func (l *Loader) GetUsers(ctx context.Context, ids ...string) ([]model.User, error) {
	thunk := l.uLoader.LoadMany(ctx, dataloader.NewKeysFromStrings(ids))
	res, err := thunk()
	if err != nil {
		return nil, err[0]
	}

	result := make([]model.User, len(res))
	for _, u := range res {
		result = append(result, *u.(*model.User))
	}

	return result, nil
}

func NewLoader(uClient users.Client, uService user.Service) *Loader {
	reader := newReader(uClient, uService)
	uLoader := dataloader.NewBatchedLoader(reader.getBatchUsers, dataloader.WithWait(wait))
	return &Loader{uLoader}
}

func ToOutgoingCtx(ctx context.Context, loader *Loader) context.Context {
	return context.WithValue(ctx, loadersKey, loader)
}

func FromIncomingCtx(ctx context.Context) *Loader {
	return ctx.Value(loadersKey).(*Loader)
}
