package service

import (
	mw "github.com/RedchilliSauce/flixcult/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//RestAPIBasePath - ..
const RestAPIBasePath = "/flixcultapi"

//DefaultMWFuncs - Commonly used middlewarefunc
var DefaultMWFuncs = []echo.MiddlewareFunc{
	middleware.Logger(),
	middleware.BasicAuth(middleware.DefaultBasicAuthConfig),
}

//SetupRestRoutes - Set up all the routes based on Routes var
func SetupRestRoutes(e *echo.Echo) {
	g := e.Group(RestAPIBasePath)
	g.Use(middleware.JWTWithConfig(mw.FlixCultJwtConfig))
	for _, r := range Routes {
		g.Match(r.Methods, r.Path, r.HandlerFunc, r.MiddlewareFuncs...)
	}
}
