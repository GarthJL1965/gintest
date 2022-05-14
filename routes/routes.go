package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/itsjamie/gin-cors"
	"github.com/gjlanc65/gintest/controllers"
	// "github.com/gjlanc65/gintest/middleware"

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

	// Docs (OpenAPI/'Swagger' (Swagger == deprecated term))
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	// Middleware
	/* NB : See imports to enable (&& go get ... )
	// Apply the middleware to the router (works with groups too)
	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: false,
		ValidateHeaders: false,
	}))

	router.NoRoute(middleware.NoRouteHandler())
	*/

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
