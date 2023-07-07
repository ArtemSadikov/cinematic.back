package token

import (
	jwt2 "cinematic.back/pkg/jwt"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Payload struct {
	Id         string
	PasswordID string
}

type Service struct {
	db              *gorm.DB
	accessTokenJwt  *jwt2.Wrapper
	refreshTokenJwt *jwt2.Wrapper
}

type AccessJWTClaims struct {
	jwt.StandardClaims
	UserId     string
	PasswordId string
}

type RefreshJWTClaims struct {
	jwt.StandardClaims
	UserId  string
	TokenId string
}

func (s *Service) ValidateRefreshToken(token string) (*RefreshJWTClaims, error) {
	claims, err := s.refreshTokenJwt.Validate(token)
	if err != nil {
		return nil, err
	}

	res := &RefreshJWTClaims{
		StandardClaims: claims.StandardClaims,
		UserId:         claims.UserId,
		TokenId:        *claims.TokenId,
	}

	return res, err
}

func (s *Service) ValidateAccessToken(token string) (*AccessJWTClaims, error) {
	claims, err := s.accessTokenJwt.Validate(token)
	if err != nil {
		return nil, err
	}

	res := &AccessJWTClaims{
		StandardClaims: claims.StandardClaims,
		UserId:         claims.UserId,
		PasswordId:     *claims.PasswordId,
	}

	return res, err
}

func (s *Service) GenerateTokens(userId, passwordId, tokenId string) (Tokens, error) {
	payload := map[string]string{
		"UserId":     userId,
		"PasswordId": passwordId,
	}
	accessToken, err := s.accessTokenJwt.GenerateToken(payload)
	if err != nil {
		return Tokens{}, err
	}

	delete(payload, "PasswordId")
	payload["TokenId"] = tokenId

	refreshToken, err := s.refreshTokenJwt.GenerateToken(payload)
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{accessToken, refreshToken}, nil
}

func NewService(
	db *gorm.DB,
	accessTokenJwt *jwt2.Wrapper,
	refreshTokenJwt *jwt2.Wrapper,
) *Service {
	return &Service{
		db:              db,
		accessTokenJwt:  accessTokenJwt,
		refreshTokenJwt: refreshTokenJwt,
	}
}
