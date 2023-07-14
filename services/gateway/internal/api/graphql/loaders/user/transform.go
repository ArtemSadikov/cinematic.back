package user

import (
	"cinematic.back/api/users/pb"
	"cinematic.back/services/gateway/internal/api/graphql/model"
)

func TransformUser(data *pb.User) *model.User {
	result := &model.User{
		ID: data.Id,
		Profile: &model.UserProfile{
			Email:    data.Profile.Email,
			Username: data.Profile.Username,
		},
		CreatedAt: data.CreatedAt.AsTime(),
		UpdatedAt: data.UpdatedAt.AsTime(),
	}

	if data.DeletedAt.IsValid() {
		deletedAt := data.DeletedAt.AsTime()
		result.DeletedAt = &deletedAt
	}

	return result
}
