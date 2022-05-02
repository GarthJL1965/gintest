package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gjlanc65/gintest/models"
	"github.com/gjlanc65/gintest/utils"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

// RFC 7807 Error Response (In Progress)
// There's a playoff here, in that in some cases you dont wish to
// reveal any info about the underlying system ...
// See also : https://github.com/mschneider82/problem
// Notes For other projects https://github.com/mschneider82/wordclouds
/*
func(c *gin.Context) {
  problem.New(
    problem.Title("houston! we have a problem"),
    problem.Status(http.StatusNotFound),
  ).WriteTo(c.Writer)
}
*/
// See also : https://github.com/MLJ-solutions/go-rfc7807
type ErrorResponse struct {
	Msg string
	//Err error
	Err string
}

var segments = []models.Segment{
	{ID: 1, LocalDate: time.Date(2022, 9, 10, 00, 00, 00, 00, time.FixedZone("UTC+10", 10*3600)), Name: "GJL", Details: "Leave Sydney", Who: "DML&GJL"},
}

// getSegments responds with the list of all segments as JSON.
func GetSegments(c *gin.Context) {
	utils.RenderContent(c, http.StatusOK, gin.H{"payload": segments})
}

// createSegment adds a segment from JSON received in the request body.
func CreateSegment(c *gin.Context) {
	var newSegment models.Segment

	// Call BindJSON to bind the received JSON to newSegment.
	//if err := c.BindJSON(&newSegment); err != nil {
	if err := c.ShouldBind(&newSegment); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": utils.Simple(verr)})
			/*
				var response = ErrorResponse{
					Msg: "Validation Error",
					Err: err,
				}
				c.AbortWithStatusJSON(http.StatusOK, response)
			*/
			return
		}

		log.Info(err)
		//log.Info().Err(err).Msg("unable to bind")
		/*
					var badParamResponse = ErrorResponse{
				Msg: "Bad Parameter",
				Err: fmt.Sprintf("id Parameter not numeric (%s)", id),
			}
			renderError(c, http.StatusBadRequest, gin.H{"error": badParamResponse})
		*/
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	// Add the new segment to the slice.
	segments = append(segments, newSegment)
	// Just show the created segment
	utils.RenderContent(c, http.StatusCreated, gin.H{"payload": newSegment})
}

// getSegmentByID locates the segment whose ID value matches the id
// parameter sent by the client, then returns that segment as a response.
func GetSegmentByID(c *gin.Context) {
	idAsStr := c.Param("id")

	if idAsInt, err := strconv.Atoi(idAsStr); err == nil {
		// Loop over the list of segments, looking for a segment
		// who's ID value matches the parameter.
		for _, segment := range segments {
			if segment.ID == idAsInt {
				utils.RenderContent(c, http.StatusOK, gin.H{"payload": segment})
				return
			}
		}
		// Handle/Respond to not found ...
		var notFoundResponse = ErrorResponse{
			Msg: "Data Not Found",
			Err: fmt.Sprintf("Segment (%d) not found", idAsInt),
		}
		utils.RenderError(c, http.StatusNotFound, gin.H{"error": notFoundResponse})
		return

	} else {
		var badParamResponse = ErrorResponse{
			Msg: "Bad Parameter",
			Err: fmt.Sprintf("id Parameter not numeric (%s)", idAsStr),
		}
		utils.RenderError(c, http.StatusBadRequest, gin.H{"error": badParamResponse})
	}
}
