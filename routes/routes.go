package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gjlanc65/gintest/controllers"

	//_ "github.com/gjlanc65/gintest/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Setup route group for the API
	// api := router.Group("/api")
	// api.GET("/segments", getSegments) -> http://localhost:8080/api/segments

	// Maybe Health ..
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Docs (OpenAPI/'Swagger' (depcrecated term))
	// docs route
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	// Real Routes
	router.GET("/segments", controllers.GetSegments)
	router.GET("/segments/:id", controllers.GetSegmentByID)
	router.POST("/segments", controllers.CreateSegment)
	/*
		router.NoRoute(func(c *gin.Context) {
			c.AbortWithStatus(http.StatusNotFound)
		})
	*/
	return router
}
