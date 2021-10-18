package constants

const (
	SwaggerAuthorizeContent = "Authorization"
)

// Testing
const (
	Port        = "3000"
	PathSwagger = "api-docs"
	RoutePrefix = "v1" // change to "/v1" to add v1 prefix
)

type BodyExample struct {
	ID string `json:"id"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
