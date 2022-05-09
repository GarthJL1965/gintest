#Go/Golang gin framework test

Testing API build using static data (for moment) and gin framework.

Extra (see utils/response.go) rendering layer to handle errors as XML,
  although as noted in controllers/segment.go, https://github.com/mschneider82/problem
  allows us the flexibility

OpenAPI (Swagger) working, needs a tidy up (duplicate models mainly, see notes below)
[GJL Note to self - Setup]
Main (main.go) imports _ "github.com/gjlanc65/gintest/docs" (from swag init)
Main (main.go) runs docs.SwaggerInfo.BasePath = "/"
Routes (routes.go) imports 	swaggerFiles "github.com/swaggo/files", ginSwagger "github.com/swaggo/gin-swagger"
Routes (routes.go) runs
```
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
```
This may be reason for duplicate models below
  swag init -d ./,./models,./controllers --parseDependency --parseInternal
Note package (model.s, controllers.) in controllers\segment.go seems better than including '.models' in init command
```
// @Success 201 {object} models.Segment
// @Failure 400 {object} controllers.ErrorResponse
```
-> http://localhost:8080/docs/index.html
-> http://localhost:8080/docs/docs.json