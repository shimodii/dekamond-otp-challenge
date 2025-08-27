package service

import (
  "time"
  "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("dekamond")

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(), // expires in 15 min
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}
