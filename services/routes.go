package myapi

import "github.com/go-martini/martini"

func InitRoutes(m *martini.ClassicMartini) {
	m.Get("/", func() string {
		return "Go to /books"
	})

	m.Group("/books", func(r martini.Router) {
		r.Get("/info", func() string {
			return "An API that allows you to work with Books"
		})
		r.Get("/:id", GetBookById)
		r.Post("/", CreateBook)
	})
}
