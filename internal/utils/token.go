package utils

import (
	"github.com/pkg/errors"
	"time"

	"github.com/UraharaKiska/go-auth/internal/model"
	"github.com/dgrijalva/jwt-go"
	// "github.com/pkg/errors"
)

func GenerateToken(info model.UserBaseInfo, secretKey []byte, duration time.Duration) (string, error) {
	claims := model.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
		Email: info.Email,
		Role: info.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString(secretKey)
}

func VerifyToken(tokenStr string, secretKey []byte) (*model.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&model.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.Errorf("unexpected token signing method")
			}
			return secretKey, nil
		},

	)
	if err != nil {
		return nil, errors.Errorf("invalid token: %s", err.Error())
	}
	claims, ok := token.Claims.(*model.UserClaims)
	if !ok {
		return nil, errors.Errorf("invalid token claims")
	}
	return claims, nil
}