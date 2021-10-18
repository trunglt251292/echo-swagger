package echoswagger

import (
	"echo-swagger/constants"

	"github.com/labstack/echo/v4"
)

// Init ...
func Init(echo *echo.Echo, info *Info) ApiRoot {
	se := New(echo, info.BasePath, info)
	se.AddSecurityAPIKey(constants.SwaggerAuthorizeContent, "Bearer token", SecurityInHeader)
	return se
}
