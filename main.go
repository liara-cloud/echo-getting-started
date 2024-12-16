package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Serve static files (CSS, JS, images)
	e.Static("/css", "static/css")
	e.Static("/js", "static/js")
	e.Static("/images", "static/images")

	// Serve the HTML file at the root
	e.File("/", "static/index.html")

	e.Logger.Fatal(e.Start(":8080")) // Start the server on port 8080
}
