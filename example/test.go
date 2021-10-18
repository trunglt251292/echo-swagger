package main

import (
	echoswagger "echo-swagger"
	"echo-swagger/constants"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Start all app handler, add more handler if need
func main() {
	e := echo.New()
	// Swagger
	swg := echoswagger.Init(e, &echoswagger.Info{
		Title:       "Swagger example",
		Description: "Exampale - API use for ....",
		Version:     "1.0.0",
		License: &echoswagger.License{
			Name: "Swagger v3",
			URL:  "https://swagger.io/docs/specification/about/",
		},
		BasePath: "/v1/api-docs", // Default
	})

	r := swg.Group("Example", constants.RoutePrefix+"/example").SetSecurity(constants.SwaggerAuthorizeContent)

	// Method Get Query
	r.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"success": true,
		})
	}).AddDescriptions(echoswagger.NewDescription("Example get api", "")).
		AddParamQuery("string", "limit", "Limit", false).
		AddParamQuery("string", "id", "ID", true)

	// Param
	r.PUT("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, echo.Map{
			"success": true,
			"id":      id,
		})
	}).AddDescriptions(echoswagger.NewDescription("Example with param", "")).
		AddParamPath("string", "id", "ID")

	// Body
	r.POST("", func(c echo.Context) error {
		var (
			body *constants.BodyExample
		)
		c.Bind(&body)
		return c.JSON(http.StatusOK, echo.Map{
			"success": true,
			"body":    body,
		})
	}).AddDescriptions(echoswagger.NewDescription("Example with param", "")).
		AddParamBody(&constants.BodyExample{ID: "1"}, "body", "Body", true).
		ResponseDefault([]*echoswagger.ResponseJSON{
			{
				Code: 200,
				Desc: "Response",
				Schema: constants.Response{
					Success: true,
					Data:    &constants.BodyExample{},
				},
				Header: nil,
			},
		})

	e.Logger.Fatal(e.Start(":" + constants.Port))
}
