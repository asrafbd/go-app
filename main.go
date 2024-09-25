package main

import (
    "net/http"
    "os"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, /airports_v2 endpoint\n")
    })

    httpPort := os.Getenv("PORT")
    if httpPort == "" {
        httpPort = "8080"
    }

    e.Logger.Fatal(e.Start(":" + httpPort))
}
