package middleware

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Credential - Struct to take hold username/password
type Credential struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

//ClaimsInp - Input to generate payload
type ClaimsInp struct {
	Name    string
	IsAdmin bool
}

var credentials = map[string]string{
	"admin":   "leviathan",
	"jonsnow": "fluffy",
}

//SigningKey -
var SigningKey = []byte("awakes") // need to get this from an .env file

//FlixCultJwtConfig -
var FlixCultJwtConfig = middleware.JWTConfig{
	SigningKey:    SigningKey,
	SigningMethod: middleware.AlgorithmHS256,
	AuthScheme:    "Bearer",
	Skipper:       SkipJWTAuthentication,
}

//BasicAuthenticate - returns (authenticationStatus, admin)
func BasicAuthenticate(username string, password string) (bool, bool) {
	authenticationStatus := false
	admin := false
	if username != "" && credentials[username] == password {
		authenticationStatus = true
		if isUserAdmin(username) {
			admin = true
		}
	}
	return authenticationStatus, admin
}

func isUserAdmin(username string) bool {
	if username == "admin" {
		return true
	}
	return false
}

//GenerateJWT - Generates a JWT token using
func GenerateJWT(claimsInp ClaimsInp) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = claimsInp.IsAdmin
	claims["name"] = claimsInp.Name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString(SigningKey)
	return tokenString
}

//SkipJWTAuthentication - Skipper function to skip JWT Auth
func SkipJWTAuthentication(c echo.Context) bool {
	skipAuth := strings.HasSuffix(c.Path(), "/login")
	skipAuth = skipAuth && c.Request().Method == "GET"
	return skipAuth
}
