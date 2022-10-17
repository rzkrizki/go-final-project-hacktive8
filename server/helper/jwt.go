package helper

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateToken(email string) (string, error) {
	err := godotenv.Load()

	if err != nil {
		return "nil", err
	}

	res, err := strconv.Atoi(os.Getenv("TIMEOUTJWT"))
	if err != nil {
		return "nil", err
	}

	timeout := time.Duration(res) * time.Hour
	secret := os.Getenv("SECRETJWT")
	payload := jwt.MapClaims{
		"email":   email,
		"expired": time.Now().Add(timeout),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signed, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func ValidateToken(tokenString string) (map[string]interface{}, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	secret := os.Getenv("SECRETJWT")

	errResp := fmt.Errorf("need signin")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResp
	}

	var payload = map[string]interface{}{}
	claims := token.Claims.(jwt.MapClaims)

	payload["email"] = claims["email"]

	exp := fmt.Sprintf("%v", claims["expired"])

	now := time.Now()
	expTime, _ := time.Parse(time.RFC3339, exp)

	if !now.Before(expTime) {
		return nil, fmt.Errorf("expired")
	}

	return payload, nil
}
