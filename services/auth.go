package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtSecretKey          = []byte("your_secret_key")         // Replace with your own secret key
	refreshTokenSecretKey = []byte("your_refresh_secret_key") // Replace with your own secret key
)

func GenerateTokens(userID uint) (string, string, error) {
	accessTokenClaims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Access token expires in 1 hour
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecretKey)
	if err != nil {
		return "", "", err
	}

	refreshTokenClaims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // Refresh token expires in 1 week
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(refreshTokenSecretKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
