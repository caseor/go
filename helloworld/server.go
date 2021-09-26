package main

import (
	serverapi "helloworld/api"
	
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", serverapi.HelloWorld)
	e.Logger.Fatal(e.Start(":8080"))
}