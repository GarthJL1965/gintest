package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gjlanc65/gintest/controllers"
)

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
	router.GET("/segments", controllers.GetSegments)
	router.GET("/segments/:id", controllers.GetSegmentByID)
	router.POST("/segments", controllers.CreateSegment)
	/*
		router.NoRoute(func(c *gin.Context) {
			c.AbortWithStatus(http.StatusNotFound)
		})
	*/
	router.Run("localhost:8080")
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
