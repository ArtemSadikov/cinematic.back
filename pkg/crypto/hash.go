package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pswd string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pswd), bcrypt.DefaultCost)

	return string(bytes)
}

func ComparePassword(pswd string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pswd))
}

func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))

	return hex.EncodeToString(hash[:])
}

func CompareToken(token, hash string) bool {
	hashed := HashToken(token)

	return hashed == hash
}
