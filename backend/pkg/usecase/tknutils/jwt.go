package tknutils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateJwt(accountId int, expireAt time.Time) (string, error) {
	jwtKey := os.Getenv("JWT_KEY")

	claims := jwt.MapClaims{
		"account_id": accountId,
		"exp":        expireAt.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return strToken, nil
}

func ParseJwt(strToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(strToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
