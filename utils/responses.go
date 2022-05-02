package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// router.AbortWithStatusJSON() exists but there's no c.AbortWithStatusXML()
// Q) Why wouldnt we send back errors as XML if the client/request system
// was XML based AND set Content-Type to application/aml ?
// .. what's the beef with this 'everything has to be JSON' ?
func RenderError(c *gin.Context, httpStatus int, data gin.H) {
	switch c.Request.Header.Get("Content-Type") {
	case "application/json":
		c.IndentedJSON(httpStatus, data["error"])
	case "application/xml":
		c.XML(httpStatus, data["error"])
	default:
		c.IndentedJSON(httpStatus, data["error"])
	}
}

func RenderContent(c *gin.Context, httpStatus int, data gin.H) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.IndentedJSON(httpStatus, data["payload"])
	case "application/xml":
		c.XML(httpStatus, data["payload"])
	default:
		plainStr := fmt.Sprint(data["payload"])
		c.String(httpStatus, plainStr)
	}
}

func Simple(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err
	}

	return errs
}
