package tokens

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
)

// const lifeTime = 12 * time.Hour

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var tokenSecret = generateSecret(32)

func GenerateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": userID,
	})

	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		log.Fatal("Couldn't sign token", err.Error())
		return "", err
	}

	return tokenString, nil
}

func DecryptToken(tokenString string) (jwt.MapClaims, error) {
	claims, err := extractClaims(tokenString)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func extractClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(tokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid JWT")
	}

}

func generateSecret(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
