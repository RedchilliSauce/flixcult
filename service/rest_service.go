package service

import (
	"github.com/labstack/echo"
)

//StartRESTService - Starts up a new REST API service
func StartRESTService(e *echo.Echo, hostname string, port string) {
	// TODO - Add any common middleware for REST
	SetupRoutes(e)
	e.Start(hostname + ":" + port)
}
