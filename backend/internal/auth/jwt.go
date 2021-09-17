package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/hansel-app/hansel/internal/core/domain/users"
)

type JWTManager struct {
	secretKey     string
	tokenDuration *time.Duration
}

func New(secretKey string, tokenDuration *time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:     secretKey,
		tokenDuration: tokenDuration,
	}
}

func (m *JWTManager) Generate(user *users.User) (string, error) {
	// An expiry claim of 0 indicates that the token never expires.
	expiresAt := int64(0)
	if m.tokenDuration != nil {
		expiresAt = time.Now().Add(*m.tokenDuration).Unix()
	}

	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		UserID: user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(m.secretKey))
}

func (m *JWTManager) Verify(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			return nil, errors.New("unexpected token signing method")
		}
		return []byte(m.secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
