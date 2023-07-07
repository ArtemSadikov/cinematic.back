package auth

import (
	"cinematic.back/pkg/crypto"
	auth2 "cinematic.back/services/users/internal/domain/models/auth"
	userModel "cinematic.back/services/users/internal/domain/models/user"
	"cinematic.back/services/users/internal/domain/models/user/profile"
	"cinematic.back/services/users/internal/domain/services"
	token2 "cinematic.back/services/users/internal/domain/services/token"
	"context"
	"errors"
	"github.com/google/uuid"
)

type UseCase struct {
	aService services.AuthService
	uService services.UserService
	tService services.TokenService
}

func (u *UseCase) RefreshToken(ctx context.Context, token string) (*token2.Tokens, error) {
	claims, err := u.tService.ValidateRefreshToken(token)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(claims.UserId)
	if err != nil {
		return nil, err
	}
	auth, err := u.aService.FindUserAuthByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}

	if ok := crypto.CompareToken(token, auth.Token.Token); !ok {
		return nil, errors.New("failed refresh compare")
	}

	tokenId := uuid.New()
	tokens, err := u.tService.GenerateTokens(
		claims.UserId,
		auth.PasswordID.String(),
		tokenId.String(),
	)
	if err != nil {
		return nil, err
	}

	auth.SetToken(tokens.RefreshToken, tokenId)

	if err := u.aService.Save(ctx, auth); err != nil {
		return nil, err
	}

	return &tokens, nil
}

func (u *UseCase) Register(ctx context.Context, email, username, password string) (*token2.Tokens, error) {
	existsUser, err := u.uService.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !existsUser.IsNull() {
		return nil, errors.New("exists")
	}

	user := userModel.User{
		Profile: profile.Profile{
			Email:    email,
			Username: username,
		},
		Auth: &auth2.UserAuth{},
	}

	user.Auth.ChangePassword(password)

	if err := u.uService.Save(ctx, &user); err != nil {
		return nil, err
	}

	tokenId := uuid.New()

	tokens, err := u.tService.GenerateTokens(user.ID.String(), user.Auth.PasswordID.String(), tokenId.String())
	if err != nil {
		return nil, err
	}

	user.Auth.SetToken(tokens.RefreshToken, tokenId)

	if err := u.aService.Save(ctx, user.Auth); err != nil {
		return nil, err
	}

	return &tokens, err
}

func (u *UseCase) AuthByCredentials(ctx context.Context, email, password string) (*token2.Tokens, error) {
	user, err := u.uService.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	auth, err := u.aService.FindUserAuthByUserID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	if err := crypto.ComparePassword(password, auth.Password); err != nil {
		return nil, err
	}

	tokenId := uuid.New()
	tokens, err := u.tService.GenerateTokens(user.ID.String(), auth.PasswordID.String(), tokenId.String())
	if err != nil {
		return nil, err
	}

	auth.SetToken(tokens.RefreshToken, tokenId)

	if err := u.aService.Save(ctx, auth); err != nil {
		return nil, err
	}

	return &tokens, nil
}

func (u *UseCase) AuthByAccessToken(ctx context.Context, token string) error {
	claims, err := u.tService.ValidateAccessToken(token)
	if err != nil {
		return err
	}

	userId, err := uuid.Parse(claims.UserId)
	if err != nil {
		return err
	}
	if _, err = u.uService.FindByID(ctx, userId); err != nil {
		return err
	}

	return nil
}

func NewUseCase(
	aService services.AuthService,
	uService services.UserService,
	tService services.TokenService,
) *UseCase {
	return &UseCase{
		aService: aService,
		uService: uService,
		tService: tService,
	}
}
