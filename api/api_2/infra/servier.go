package infra

import (
	util "module/util"

	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	SERVER_PORT      string
	SELF_ECHO_SERVER *echo.Echo
)

func init() {
	server_port := util.GetEnv("SERVER_PORT", "8080")
	SERVER_PORT = fmt.Sprintf(":%s", server_port)
}

func FactoryServer() *echo.Echo {
	if SELF_ECHO_SERVER == nil {
		SELF_ECHO_SERVER = echo.New()
	}
	var e *echo.Echo = SELF_ECHO_SERVER

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	return SELF_ECHO_SERVER
}
