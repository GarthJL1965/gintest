package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gjlanc65/gintest/utils"
)

type ErrorResponse struct {
	Msg string
	//Err error
	Err string
}

// NoMethodHandler
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var noMethodHandlerResponse = ErrorResponse{
			Msg: "NoMethod",
			Err: "Method not found/permitted",
		}
		utils.RenderError(c, http.StatusMethodNotAllowed, gin.H{"error": noMethodHandlerResponse})
	}
}

// NoRouteHandler
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var noRouteHandlerResponse = ErrorResponse{
			Msg: "NoRoute",
			Err: "No matching route found",
		}
		utils.RenderError(c, http.StatusNotFound, gin.H{"error": noRouteHandlerResponse})
	}
}
