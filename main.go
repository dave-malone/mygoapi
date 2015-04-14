package main

import (
	myapi "github.com/dave-malone/myapi/services"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	myapi.InitDb()

	m := martini.Classic()
	m.Use(render.Renderer())

	myapi.InitRoutes(m)

	m.Run()
}
