package main

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	users = []string{"Joe", "Veer", "Zion"}
)

func getUsers(e echo.Context) error {
	return e.JSON(http.StatusOK, users)
}

// allowOrigin takes the origin as an argument and returns true if the origin
// is allowed or false otherwise.
func allowOrigin(origin string) (bool, error) {
	return regexp.MatchString(`^https:\/\/labstack\.(net|com)$`, origin)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS default
	// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	// e.Use(middleware.CORS())

	// CORS restricted
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: allowOrigin,
		AllowMethods:    []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/api/users", getUsers)
	
	e.Logger.Fatal(e.Start(":1323"))

}
