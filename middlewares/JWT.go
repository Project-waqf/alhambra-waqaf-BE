package middlewares

import (
	"os"
	"strconv"
	"time"
	"wakaf/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: middleware.AlgorithmHS256,
		SigningKey:    []byte(config.Getconfig().SECRET_JWT),
	})
}

func CreateToken(id int, email string) (string, error) {
	sec := os.Getenv("JWT_EXPIRY")
	duration, err := strconv.Atoi(sec)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["IdUser"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Second * time.Duration(duration)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Getconfig().SECRET_JWT))
}
