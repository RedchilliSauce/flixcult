package service

import "github.com/labstack/echo"
import "github.com/labstack/echo/middleware"

//DefaultMWFuncs - Commonly used middlewarefunc
var DefaultMWFuncs = []echo.MiddlewareFunc{
	middleware.Logger(),
	middleware.BasicAuth(BasicAuth),
}

//BasicAuth function
func BasicAuth(username string, password string, c echo.Context) bool {
	if username == "flix" && password == "cult" {
		return true
	}
	return false
}

//SetupRoutes - Set up all the routes based on Routes var
func SetupRoutes(e *echo.Echo) {
	for _, r := range Routes {
		e.Match(r.Methods, r.Path, r.HandlerFunc, r.MiddlewareFuncs...)
	}
}
