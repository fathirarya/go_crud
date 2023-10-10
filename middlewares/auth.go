package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id uint, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET")))
}
