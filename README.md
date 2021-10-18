# Echo swagger


## Getting started
1. Download echo-swagger by using:
```sh
$ go get -u github.com/trunglt251292/echo-swagger
```
To build from source you need Go v4

## How to use

Find the example source code [here](https://github.com/trunglt251292/echo-swagger/tree/main/example).

1. Init echo swagger code:

```go
type (
    Info struct {
        Title          string                
        Description    string                 
        TermsOfService string                 
        Contact        *Contact               
        License        *License               
        Version        string                 
        Host           string                 
        BasePath       string                 
    }
)
```

```go
package main

import (
  echoswagger "echo-swagger"
  "github.com/labstack/echo/v4"
)
func main() {
	// Init echo
	e := echo.New()
    // Swagger
    swg := echoswagger.Init(e, &echoswagger.Info{
        Title:       "Swagger example", // Title swagger
        Description: "Exampale - API use for ....", // Desc swagger
        Version:     "1.0.0", // Version api
        License: &echoswagger.License{
            Name: "Swagger v3",
            URL:  "https://swagger.io/docs/specification/about/",
        }, // License
        BasePath: "/v1/api-docs", // Default
    })
}
```

2. Example some API. Example using echo: [here](https://github.com/trunglt251292/echo-swagger/tree/main/example).

3. Run your `make run-test`:

![example](https://github.com/trunglt251292/echo-swagger/blob/main/assets/example.png?raw=true)