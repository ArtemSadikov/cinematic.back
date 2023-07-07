package grpc

import (
	users "cinematic.back/services/users/internal/delivery/grpc/interface/users"
	"cinematic.back/services/users/internal/delivery/grpc/mappers"
	"context"
	"github.com/google/uuid"
)

func (s *Server) CreateUser(ctx context.Context, in *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	model := mappers.MakeModelFromWriteData(in.Data)

	user, err := s.usersUC.Create(ctx, &model)
	if err != nil {
		return nil, err
	}

	resp := &users.CreateUserResponse{
		User: mappers.MakeUserFromModel(*user),
	}

	return resp, nil
}

func (s *Server) UpdateUser(ctx context.Context, in *users.UpdateUserByIdRequest) (*users.UpdateUserByIdResponse, error) {
	model, err := mappers.MakeModelFromWriteDataWithId(in.Data, in.Id)
	if err != nil {
		return nil, err
	}

	user, err := s.usersUC.UpdateByID(ctx, model.ID, &model)
	if err != nil {
		return nil, err
	}

	resp := &users.UpdateUserByIdResponse{
		User: mappers.MakeUserFromModel(*user),
	}

	return resp, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *users.DeleteUserByIdRequest) (*users.DeleteUserByIdResponse, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}

	model, err := s.usersUC.DeleteUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &users.DeleteUserByIdResponse{
		User: mappers.MakeUserFromModel(*model),
	}

	return resp, nil
}

func (s *Server) FindUserById(ctx context.Context, in *users.FindUserByIdRequest) (*users.FindUserByIdResponse, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}

	model, err := s.usersUC.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &users.FindUserByIdResponse{
		User: mappers.MakeUserFromModel(*model),
	}

	return resp, nil
}
