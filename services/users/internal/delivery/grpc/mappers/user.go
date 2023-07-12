package mappers

import (
	"cinematic.back/api/users/pb"
	"cinematic.back/services/users/internal/domain/models/user"
	"cinematic.back/services/users/internal/domain/models/user/profile"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MakeModelFromWriteData(data *pb.UserWrite) user.User {
	return user.User{
		Profile: profile.Profile{
			Email:    data.Email,
			Username: data.Username,
		},
	}
}

func MakeModelFromWriteDataWithId(data *pb.UserWrite, id string) (user.User, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return user.User{}, err
	}

	model := MakeModelFromWriteData(data)
	model.ID = parsedId

	return model, nil
}

func MakeUserFromModel(model user.User) *pb.User {
	result := pb.User{
		Id: model.ID.String(),
		Profile: &pb.UserProfile{
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
