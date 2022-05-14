module gintest

go 1.17

require (
	github.com/gjlanc65/gintest/routes v0.0.0
	//github.com/gjlanc65/gintest/docs
	github.com/swaggo/swag v1.8.1 // indirect
)

require github.com/gjlanc65/gintest/docs v0.0.0

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.7.7 // indirect
	github.com/gjlanc65/gintest/controllers v0.0.0 // indirect
	github.com/gjlanc65/gintest/db v0.0.0 // indirect
	github.com/gjlanc65/gintest/models v0.0.0 // indirect
	github.com/gjlanc65/gintest/utils v0.0.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.6 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2 // indirect
	github.com/swaggo/gin-swagger v1.4.3 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20220507011949-2cf3adece122 // indirect
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20220503163025-988cb79eb6c6 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.10 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/gjlanc65/gintest/controllers => ./controllers
	github.com/gjlanc65/gintest/db => ./db
	github.com/gjlanc65/gintest/docs => ./docs
	github.com/gjlanc65/gintest/mmiddleware => ./mmiddleware
	github.com/gjlanc65/gintest/models => ./models
	github.com/gjlanc65/gintest/routes => ./routes
	github.com/gjlanc65/gintest/utils => ./utils
)
