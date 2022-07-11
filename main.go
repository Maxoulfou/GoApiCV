package main

import (
	"net/http"

	_ "github.com/Devansh3712/portfolio/docs"
	"github.com/joho/godotenv"

	"github.com/Devansh3712/portfolio/errors"
	"github.com/Devansh3712/portfolio/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Portfolio API
// @version 1.0
// @description Devansh's portfolio as a REST API.

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @BasePath /
func main() {
	// Load .env file
	godotenv.Load(".env")

	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	app.RedirectTrailingSlash = true
	app.HandleMethodNotAllowed = true

	app.NoRoute(func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusNotFound, errors.EndpointNotFound)
	})
	app.NoMethod(func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusMethodNotAllowed, errors.MethodNotAllowed)
	})

	app.GET("/", routes.Root)
	app.GET("/about", routes.About)
	app.GET("/projects", routes.Projects)
	app.GET("/misc", routes.Misc)
	app.POST("/contact", routes.Contact)

	app.StaticFile("/resume", "./resources/resume.pdf")
	// Documentation route => /docs/index.html
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Run(":9999")
}
