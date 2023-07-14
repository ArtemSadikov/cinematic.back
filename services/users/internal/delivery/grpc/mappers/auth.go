package mappers

import (
	"cinematic.back/api/users/pb"
	"cinematic.back/services/users/internal/domain/services/token"
)

func MakeTokens(tokens *token.Tokens) *pb.Tokens {
	return &pb.Tokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
}
