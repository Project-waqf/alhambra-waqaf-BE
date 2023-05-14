package main

import (
	"context"
	"fmt"
	"wakaf/config"
	"wakaf/factory"
	"wakaf/pkg/helper"
	"wakaf/utils/database/mysql"
	"wakaf/utils/database/redis"

	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.Getconfig()
	db := mysql.InitDBmysql(config)
	redis := redis.InitRedis(config)

	var log = helper.Logger()
	log.Info("Connecting Redis", zap.Any("PING", redis.Ping(context.Background())))
	log.Info("Connecting MYSQL", zap.Any("Error", db.Error))

	helper.InitMigrate(db)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))


	factory.InitFactory(e, db, &redis, &log)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.SERVER_PORT)))
}