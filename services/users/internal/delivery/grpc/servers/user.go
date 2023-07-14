package servers

import (
	"cinematic.back/api/users/pb"
	"cinematic.back/services/users/internal/delivery/grpc/mappers"
	"cinematic.back/services/users/internal/usecase"
	"context"
	"github.com/google/uuid"
)

type UserServer struct {
	usersUC usecase.UsersUseCases
}

func (s *UserServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	model := mappers.MakeModelFromWriteData(in.Data)

	user, err := s.usersUC.Create(ctx, &model)
	if err != nil {
		return nil, err
	}

	resp := &pb.CreateUserResponse{
		User: mappers.MakeUserFromModel(*user),
	}

	return resp, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, in *pb.UpdateUserByIdRequest) (*pb.UpdateUserByIdResponse, error) {
	model, err := mappers.MakeModelFromWriteDataWithId(in.Data, in.Id)
	if err != nil {
		return nil, err
	}

	user, err := s.usersUC.UpdateByID(ctx, model.ID, &model)
	if err != nil {
		return nil, err
	}

	resp := &pb.UpdateUserByIdResponse{
		User: mappers.MakeUserFromModel(*user),
	}

	return resp, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, in *pb.DeleteUserByIdRequest) (*pb.DeleteUserByIdResponse, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}

	model, err := s.usersUC.DeleteUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &pb.DeleteUserByIdResponse{
		User: mappers.MakeUserFromModel(*model),
	}

	return resp, nil
}

func (s *UserServer) FindUserById(ctx context.Context, in *pb.FindUserByIdRequest) (*pb.FindUserByIdResponse, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}

	model, err := s.usersUC.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &pb.FindUserByIdResponse{
		User: mappers.MakeUserFromModel(*model),
	}

	return resp, nil
}

func (s *UserServer) FindUsersByIds(ctx context.Context, in *pb.FindUsersByIdsRequest) (*pb.FindUsersByIdsResponse, error) {
	var ids []uuid.UUID

	for _, id := range in.Ids {
		parsed, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}

		ids = append(ids, parsed)
	}

	res, err := s.usersUC.FindUsersByIds(ctx, ids...)
	if err != nil {
		return nil, err
	}

	resp := &pb.FindUsersByIdsResponse{Users: []*pb.User{}}

	for _, user := range res {
		resp.Users = append(resp.Users, mappers.MakeUserFromModel(*user))
	}

	return resp, nil
}

func NewUserServer(usersUC usecase.UsersUseCases) *UserServer {
	return &UserServer{usersUC}
}
