package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gjlanc65/gintest/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Setup route group for the API
	// api := router.Group("/api")
	// api.GET("/segments", getSegments) -> http://localhost:8080/api/segments

	// Maybe Health ..
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello From gintest")
	})
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
