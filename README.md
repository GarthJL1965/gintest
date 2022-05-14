# Go/Golang gin framework test

Testing API build using static data (for moment) and gin framework.

Extra (see utils/response.go) rendering layer to handle errors as XML,
  although as noted in controllers/segment.go, https://github.com/mschneider82/problem
  allows us the flexibility

OpenAPI (Swagger) working, needs a tidy up (duplicate models mainly, see notes below)<br>

## [GJL Note to self - OpenAPI Setup]

* go get -u github.com/swaggo/swag/cmd/swag
* go install github.com/swaggo/swag/cmd/swag
* go get -u github.com/swaggo/gin-swagger
* go get -u github.com/swaggo/files

Main (main.go) imports _ "github.com/gjlanc65/gintest/docs" (from swag init)<br>

Main (main.go) runs docs.SwaggerInfo.BasePath = "/"<br>

Routes (routes.go) imports 	swaggerFiles "github.com/swaggo/files", ginSwagger
"github.com/swaggo/gin-swagger"<br>

Routes (routes.go) runs
```
router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
ginSwagger.WrapHandler(swaggerFiles.Handler,
	ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
	ginSwagger.DefaultModelsExpandDepth(-1))
```
This may be reason for duplicate models below<br>

```plaintext
swag init -d ./,./models,./controllers --parseDependency --parseInternal
```

Note package (models., controllers.) in controllers\segment.go seems better than including '.models' in init command
```
// @Success 201 {object} models.Segment
// @Failure 400 {object} controllers.ErrorResponse
```

## API Doc Routes are therefore :-

-> http://localhost:8080/docs/index.html<br>
-> http://localhost:8080/docs/docs.json

## TODO :
* Fix controllers\segment.go error handling in CreateSegment Code (JSON|XML)

## Test data for Postman
```JSON
{
    "id": 3,
    "LocalDate": "2022-09-11T17:10:00+10:00",
    "name": "RMN",
    "details": "Depart",
    "who": "RMN&&MMN"
}
```
```XML
<segment>
    <ID>1</ID>
    <LocalDate>2022-09-10T00:00:00+10:00</LocalDate>
    <Name>RMN</Name>
    <Details>Depart</Details>
    <Who>RMN&amp;&amp;MMN</Who>
</segment>
```

## Acknowledgments

See : https://www.stackhawk.com/blog/configuring-cors-for-go/ for nice CORS pic

WestWind 'Markdown Monster' https://markdownmonster.west-wind.com/ used for Markdown