package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	//e.Pre(middleware.RemoveTrailingSlash())
	//e.Use(middleware.RequestID())
	//e.Use(middleware.Gzip())
	//e.Use(middleware.Recover())
	//e.Use(middleware.Secure())

	v1 := e.Group("/v1")
	groupV1(v1)
	groupV1Users(v1)

	e.Logger.Fatal(e.Start(":1323"))
}

func groupV1(e *echo.Group) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK GET / (not a group)")
	})
	e.GET("/users/:id/profile", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK GET /users/:id/profile (not a group)")
	})
	e.DELETE("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK DELETE /users/:id (not a group)")
	})
}

func groupV1Users(e *echo.Group) {
	ug := e.Group("/users")
	ug.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK GET /users")
	})
	ug.GET("/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK GET /users/:id")
	})
}
