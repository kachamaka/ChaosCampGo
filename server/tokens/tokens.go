package tokens

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

// const lifeTime = 12 * time.Hour

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateToken is a function that generates JWT from unique user ID
func GenerateToken(userID string) (string, error) {
	log.Println(viper.GetString("TOKEN_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": userID,
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("TOKEN_SECRET")))
	if err != nil {
		log.Fatal("Couldn't sign token", err.Error())
		return "", err
	}

	return tokenString, nil
}

// DecryptToken is a function that decrypts JWT and returns JWT claims
func DecryptToken(tokenString string) (jwt.MapClaims, error) {
	claims, err := extractClaims(tokenString)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// extractClaims is a function that extracts JWT claims from JWT
func extractClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(viper.GetString("TOKEN_SECRET")), nil
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

// GenerateSecret is a function that generates random sequence of letters and numbers
func GenerateSecret(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
