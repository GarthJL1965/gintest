package main

import (
	"github.com/gjlanc65/gintest/docs"
	_ "github.com/gjlanc65/gintest/docs"
	"github.com/gjlanc65/gintest/routes"
)

// @title Go + Gin Trip Segment API
// @version 1.0
// @description This is a sample server. You can visit the GitHub repository at https://github.com/GJLANC65/gintest

// @contact.name API Support
// @contact.url https://github.com/GJLANC65/gintest
// @contact.email [protected]

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

func main() {
	docs.SwaggerInfo.BasePath = "/"

	router := routes.SetupRouter()
	router.Run("localhost:8080")
}
