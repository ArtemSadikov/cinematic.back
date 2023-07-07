package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type Wrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type claims struct {
	jwt.StandardClaims
	UserId     string  `json:"userId"`
	PasswordId *string `json:"passwordId,omitempty"`
	TokenId    *string `json:"tokenId,omitempty"`
}

func (s *Wrapper) GenerateToken(payload map[string]string) (string, error) {
	claims := &claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(s.ExpirationHours)).Unix(),
			Issuer:    s.Issuer,
		},
		UserId: payload["UserId"],
	}

	if v, ok := payload["PasswordId"]; ok {
		claims.PasswordId = &v
	}

	if v, ok := payload["TokenId"]; ok {
		claims.TokenId = &v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *Wrapper) Validate(data string) (*claims, error) {
	token, err := jwt.ParseWithClaims(
		data,
		&claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.SecretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*claims)

	if !ok {
		return nil, errors.New("Couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil
}

func NewJwtWrapper(
	secretKey string,
	issuer string,
	expirationHours int64,
) *Wrapper {
	return &Wrapper{
		SecretKey:       secretKey,
		Issuer:          issuer,
		ExpirationHours: expirationHours,
	}
}
