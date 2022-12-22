package middlewares

import (
	"time"
	"wakaf/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: middleware.AlgorithmHS256,
		SigningKey: []byte(config.Getconfig().SECRET_JWT),
	})
}

func CreateToken(id int, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] =  true
	claims["IdUser"] = id
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Getconfig().SECRET_JWT))
}