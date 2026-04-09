package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidToken = errors.New("invalid token")

var sharedSecret []byte

func InitSecret(secret string) {
	sharedSecret = []byte(secret)
}

func IssueToken(userID, email, role string) (string, error) {
	return GenerateToken(sharedSecret, userID, email, role, 24*time.Hour)
}

func ParseClaimsFromRequest(tokenString string) (jwt.MapClaims, error) {
	return ParseToken(sharedSecret, tokenString)
}

func GenerateToken(secret []byte, userID, email, role string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(ttl).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(secret)
}

func ParseToken(secret []byte, tokenString string) (jwt.MapClaims, error) {
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return secret, nil
	})
	if err != nil || !t.Valid {
		return nil, ErrInvalidToken
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
