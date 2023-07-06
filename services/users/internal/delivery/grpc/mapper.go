package grpc

import (
	users "cinematic.back/services/users/internal/delivery/grpc/interface"
	"cinematic.back/services/users/internal/domain/models/user"
	"cinematic.back/services/users/internal/domain/models/user/profile"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func makeModelFromWriteData(data *users.UserWrite) user.User {
	return user.User{
		Profile: profile.Profile{
			Email:    data.Email,
			Username: data.Username,
		},
	}
}

func makeModelFromWriteDataWithId(data *users.UserWrite, id string) (user.User, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return user.User{}, err
	}

	model := makeModelFromWriteData(data)
	model.ID = parsedId

	return model, nil
}

func makeUserFromModel(model user.User) *users.User {
	result := users.User{
		Id: model.ID.String(),
		Profile: &users.UserProfile{
			Username: model.Profile.Username,
			Email:    model.Profile.Email,
		},
		CreatedAt: timestamppb.New(model.CreatedAt),
		UpdatedAt: timestamppb.New(model.UpdatedAt),
	}

	if model.DeletedAt.Valid {
		result.DeletedAt = timestamppb.New(model.DeletedAt.Time)
	}

	return &result
}
