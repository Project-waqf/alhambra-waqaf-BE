package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
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

	ctx, cancel := context.WithCancel(context.Background())

	serverErr := make(chan os.Signal, 1)
	signal.Notify(serverErr, os.Interrupt)

	factory.InitFactory(e, db, &redis, &log)
	go func() {
		e.Logger.Info("Server started")
		if err := e.Start(fmt.Sprintf(":%v", config.SERVER_PORT)); err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down server")
		}
	}()

	select {
	case <-serverErr:
		e.Logger.Print("Shutting down server gracefully...")

		shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancelShutdown()

		if err := e.Shutdown(shutdownCtx); err != nil {
			e.Logger.Printf("Server shutdown error: %v", err)
		}
		e.Logger.Info("Server gracefully stopped")
		cancel()
	case <- ctx.Done():
		e.Logger.Info("Server stopped")
	}

}
