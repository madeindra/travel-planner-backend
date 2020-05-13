package helpers

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(email string, timenow time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["iat"] = timenow
	claims["exp"] = timenow.Add(time.Hour * 72).Unix()

	t, _ := token.SignedString([]byte("MySecretKey"))

	return t, nil
}

func ValidateToken(tokenString string) error {
	var err error

	if tokenString == "" {
		return errors.New("Invalid token")
	}

	splitToken := strings.Split(tokenString, "Bearer ")
	if len(splitToken) != 2 {
		return errors.New("Invalid token")
	}

	_, err = jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		return []byte("MySecretKey"), nil
	})

	return err
}
