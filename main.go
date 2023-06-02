package main

import (
	"github.com/render-examples/go-gin-web-server/controller"
)

func main() {
	r := controller.StartWebServer()
	r.Run()
}
