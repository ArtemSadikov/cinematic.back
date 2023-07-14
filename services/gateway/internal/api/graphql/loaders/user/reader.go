package user

import (
	"cinematic.back/api/users"
	"cinematic.back/pkg/utils"
	"cinematic.back/services/gateway/internal/api/graphql/model"
	"cinematic.back/services/gateway/internal/services/user"
	"context"
	"errors"
	"github.com/graph-gophers/dataloader"
)

type reader struct {
	uClient  users.Client
	uService user.Service
}

func (r *reader) getBatchUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	res, err := r.uClient.FindUsersByIds(ctx, utils.NewStringsFromKeys(keys)...)
	if err != nil {
		return nil
	}

	mp := make(map[string]*model.User)
	for _, u := range res.Users {
		mp[u.Id] = TransformUser(u)
	}

	output := make([]*dataloader.Result, len(res.Users))

	for _, userKey := range keys {
		u, ok := mp[userKey.String()]
		res := &dataloader.Result{
			Data:  nil,
			Error: nil,
		}
		if ok {
			res.Data = u
		} else {
			res.Error = errors.New("u not found")
		}
		output = append(output, res)
	}

	return output
}

func newReader(
	uClient users.Client,
	uService user.Service,
) *reader {
	return &reader{uClient, uService}
}
