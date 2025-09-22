package jwthelper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/qulDev/jwt-gin-gorm/internal/config"
)

type Claims struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(id, role string) (string, error) {
	secret := config.GetJWTSecret()

	claims := Claims{
		ID:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   id,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	// take secret key
	secret := config.GetJWTSecret()

	// make container for claims
	claims := &Claims{}

	// parse token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		// validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// return secret key if method is valid
		return []byte(secret), nil
	})

	// check for parsing error
	if err != nil {
		return nil, err
	}

	// check if token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	//	 check expiration
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("token is expired")
	}

	// if everything is good, return claims
	return claims, nil
}
