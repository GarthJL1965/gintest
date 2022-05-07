module gintest

go 1.17

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/gjlanc65/gintest/controllers v0.0.0
	github.com/gjlanc65/gintest/models v0.0.0
	github.com/gjlanc65/gintest/db v0.0.0
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gjlanc65/gintest/utils v0.0.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20220427172511-eb4f295cb31f // indirect
	golang.org/x/sys v0.0.0-20220429233432-b5fbb4746d32 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/gjlanc65/gintest/controllers => ./controllers
	github.com/gjlanc65/gintest/models => ./models
	github.com/gjlanc65/gintest/utils => ./utils
	github.com/gjlanc65/gintest/db => ./db
)
