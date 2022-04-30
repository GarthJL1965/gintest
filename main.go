package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

// segment represents data about a trip segment
type segment struct {
	ID        int       `json:"id"`
	LocalDate time.Time `form:"local_date" binding:"required" time_format:"2006-01-02"`
	Name      string    `json:"name"`
	Details   string    `json:"details"`
	Who       string    `json:"who"`
}

var segments = []segment{
	{ID: 1, LocalDate: time.Date(2022, 9, 10, 00, 00, 00, 00, time.FixedZone("UTC+10", 10*3600)), Name: "GJL", Details: "Leave Sydney", Who: "DML&GJL"},
}

// getSegments responds with the list of all segments as JSON.
func getSegments(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, segments)
}

// postSegments adds a segment from JSON received in the request body.
func postSegments(c *gin.Context) {
	var newSegment segment

	// Call BindJSON to bind the received JSON to newSegment.
	if err := c.BindJSON(&newSegment); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": Simple(verr)})
			return
		}

		log.Info(err)
		//log.Info().Err(err).Msg("unable to bind")
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	// Add the new segment to the slice.
	segments = append(segments, newSegment)
	c.IndentedJSON(http.StatusCreated, newSegment)
}

// getSegmentByID locates the segment whose ID value matches the id
// parameter sent by the client, then returns that segment as a response.
func getSegmentByID(c *gin.Context) {
	id := c.Param("id")

	// Should have better validation here
	tmp_id := 0
	if idAsInt, err := strconv.Atoi(id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad id parameter"})
		return
	} else {
		tmp_id = idAsInt
	}
	// Loop over the list of segments, looking for a segment
	// who's ID value matches the parameter.
	for _, a := range segments {
		if a.ID == tmp_id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "segment not found"})
}

func main() {
	router := gin.Default()
	router.GET("/segments", getSegments)
	router.GET("/segments/:id", getSegmentByID)
	router.POST("/segments", postSegments)

	router.Run("localhost:8080")
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

// https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/?msclkid=968ca1ffc84411eca850456bea9df49b

/*
POST data for postman

    {
        "id": 3,
        "LocalDate": "2022-09-11T17:10:00+10:00",
        "name": "GJL",
        "details": "Arrive Rome",
        "who": "DML&GJL"
    }
*/
