package grpc

import (
	users "cinematic.back/services/users/internal/delivery/grpc/interface"
	"context"
	"github.com/google/uuid"
)

func (s *Server) CreateUser(ctx context.Context, in *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	model := makeModelFromWriteData(in.Data)

	user, err := s.uUseCase.Create(ctx, &model)
	if err != nil {
		return nil, err
	}

	resp := &users.CreateUserResponse{
		User: makeUserFromModel(*user),
	}

	return resp, nil
}

func (s *Server) UpdateUser(ctx context.Context, in *users.UpdateUserByIdRequest) (*users.UpdateUserByIdResponse, error) {
	model, err := makeModelFromWriteDataWithId(in.Data, in.Id)
	if err != nil {
		return nil, err
	}

	user, err := s.uUseCase.UpdateByID(ctx, model.ID, &model)
	if err != nil {
		return nil, err
	}

	resp := &users.UpdateUserByIdResponse{
		User: makeUserFromModel(*user),
	}

	return resp, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *users.DeleteUserByIdRequest) (*users.DeleteUserByIdResponse, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}

	model, err := s.uUseCase.DeleteUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &users.DeleteUserByIdResponse{
		User: makeUserFromModel(*model),
	}

	return resp, nil
}

func (s *Server) FindUserById(ctx context.Context, in *users.FindUserByIdRequest) (*users.FindUserByIdResponse, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}

	model, err := s.uUseCase.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &users.FindUserByIdResponse{
		User: makeUserFromModel(*model),
	}

	return resp, nil
}
