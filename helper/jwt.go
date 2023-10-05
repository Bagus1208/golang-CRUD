package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(signKey string, refreshKey string, userID int) map[string]any {
	var result = make(map[string]any)
	var accessToken = generateToken(signKey, userID)
	if accessToken == "" {
		return nil
	}

	var refreshToken = generateRefreshToken(refreshKey, accessToken)
	if refreshToken == "" {
		return nil
	}

	result["access_token"] = accessToken
	result["refresh_token"] = refreshToken

	return result
}

func generateToken(signKey string, userID int) string {
	var claims = jwt.MapClaims{}
	claims["id"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := sign.SignedString([]byte(signKey))

	if err != nil {
		return ""
	}

	return validToken
}

func generateRefreshToken(signKey string, accessToken string) string {
	var claims = jwt.MapClaims{}
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := sign.SignedString([]byte(signKey))

	if err != nil {
		return ""
	}

	return refreshToken
}
