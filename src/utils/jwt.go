package utils

import (
	"errors"
	"os"
	"time"

	"github.com/Sahil2k07/rms-go/src/database"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id int, email string, userType database.UserType) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("JWT secret not found in environment variables")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"email":    email,
		"userType": userType,
		"exp":      time.Now().Add(72 * time.Hour).Unix(), // Token expiration (3 days)
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (int, string, string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return 0, "", "", errors.New("JWT secret not found in environment variables")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return 0, "", "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return 0, "", "", errors.New("token has expired")
			}
		}

		idFloat, idOk := claims["id"].(float64)
		email, emailOk := claims["email"].(string)
		userType, userTypeOk := claims["userType"].(string)

		if !idOk || !emailOk || !userTypeOk {
			return 0, "", "", errors.New("invalid token claims")
		}

		id := int(idFloat)
		return id, email, userType, nil
	}

	return 0, "", "", errors.New("invalid token")
}
