package main

import (
	// "fmt"
	"fmt"
	"wakaf/config"
	"wakaf/factory"
	"wakaf/pkg/helper"
	"wakaf/utils/database/mysql"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.Getconfig()
	db := mysql.InitDBmysql(config)

	helper.InitMigrate(db)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))


	factory.InitFactory(e, db)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.SERVER_PORT)))
}