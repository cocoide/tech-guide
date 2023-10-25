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

type CustomClaims struct {
	AccountID int     `json:"account_id"`
	Exp       float64 `json:"exp"`
}

func ParseJwt(strToken string) (*CustomClaims, error) {
	response := &CustomClaims{}
	token, err := jwt.Parse(strToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		response.AccountID = int(claims["account_id"].(float64))
		response.Exp = claims["exp"].(float64)
	} else {
		return nil, fmt.Errorf("Failed to get user data form claims")
	}
	return response, nil
}
