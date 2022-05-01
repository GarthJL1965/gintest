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
	renderContent(c, http.StatusOK, gin.H{"payload": segments})
}

// createSegment adds a segment from JSON received in the request body.
func createSegment(c *gin.Context) {
	var newSegment segment

	// Call BindJSON to bind the received JSON to newSegment.
	//if err := c.BindJSON(&newSegment); err != nil {
	if err := c.ShouldBind(&newSegment); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": Simple(verr)})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	// Add the new segment to the slice.
	segments = append(segments, newSegment)
	//renderContent(c, http.StatusCreated, gin.H{"payload": segments})
	// Just show the created segment
	renderContent(c, http.StatusCreated, gin.H{"payload": newSegment})
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
	for _, segment := range segments {
		if segment.ID == tmp_id {
			//c.IndentedJSON(http.StatusOK, a)
			renderContent(c, http.StatusOK, gin.H{"payload": segment})
			return
		}
	}
	//c.IndentedJSON(http.StatusNotFound, gin.H{"message": "segment not found"})
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "segment not found"})
}

func main() {
	router := gin.Default()

	// Setup route group for the API
	// api := router.Group("/api")
	// api.GET("/segments", getSegments) -> http://localhost:8080/api/segments

	// Maybe Health ..
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello From gintest")
	})
	// Real Routes
	router.GET("/segments", getSegments)
	router.GET("/segments/:id", getSegmentByID)
	router.POST("/segments", createSegment)

	router.Run("localhost:8080")
}

func renderContent(c *gin.Context, httpStatus int, data gin.H) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.IndentedJSON(httpStatus, data["payload"])
		//c.IndentedJSON(http.StatusOK, data["payload"])
		//c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(httpStatus, data["payload"])
	//case "application/html":
	// If doing this need template String as param to renderContent
	//	c.HTML(http.StatusOK, "template", data["payload"])
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

// https://dev.to/codehakase/building-a-web-app-with-go-gin-and-react-5ke?msclkid=d1dde390c85611eca4a763ef3bedf3dc
// https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/?msclkid=968ca1ffc84411eca850456bea9df49b
// https://medium.com/@ekosuprastyo15/gin-gonic-mysql-golang-example-9185f202e968
// https://github.com/EDDYCJY/go-gin-example?msclkid=4b72c856c85311ec90f35187eb276ac9

/*
POST data for postman

    {
        "id": 3,
        "LocalDate": "2022-09-11T17:10:00+10:00",
        "name": "GJL",
        "details": "Arrive Rome",
        "who": "DML&GJL"
    }

	<segment>
		<ID>1</ID>
		<LocalDate>2022-09-10T00:00:00+10:00</LocalDate>
		<Name>GJL</Name>
		<Details>Leave Sydney</Details>
		<Who>DML&amp;GJL</Who>
	</segment>

*/

/*
	// gin.H is a shortcut for map[string]interface{}
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// The specific definition of protobuf is written in the testdata/protoexample file.
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// Note that data becomes binary data in the response
		// Will output protoexample.Test protobuf serialized data
		c.ProtoBuf(http.StatusOK, data)
	})
*/
