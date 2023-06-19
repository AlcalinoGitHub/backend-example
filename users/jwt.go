package users

import (
	"backend/models"
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func createJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"Username": user.Username,
		"Password": user.Password,
		"Pfp":      user.Pfp,
		"ID":       user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret_key := os.Getenv("JWT_SECRET_KEY")

	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to extract claims")
	}

	return claims, nil
}
