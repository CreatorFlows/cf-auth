package utils

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/creatorflows/cf-auth/config"
	"github.com/creatorflows/cf-auth/database/models"
)

func CreateClaims(role, email string, exp_time int64) (string, error) {
	claims := &models.Claims{
		Role: role,
		StandardClaims: jwt.StandardClaims{
			Subject:   email,
			ExpiresAt: exp_time,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JWT_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (claims *models.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)

	if !ok {
		return nil, err
	}

	return claims, nil
}
