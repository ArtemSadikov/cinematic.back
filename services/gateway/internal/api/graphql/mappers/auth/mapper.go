package auth

import (
	"cinematic.back/api/users/pb"
	"cinematic.back/services/gateway/internal/api/graphql/model"
)

func MakeTokens(tokens *pb.Tokens) *model.TokensResponse {
	return &model.TokensResponse{
		AccessToken:  tokens.GetAccessToken(),
		RefreshToken: tokens.GetRefreshToken(),
	}
}
