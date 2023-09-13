package router

import (
	"fiber_postgres/service"

	"github.com/gofiber/fiber/v2"
)

func AppSetup(app *fiber.App) {
	v1 := app.Group("/api/v1")
	BookRouter(v1)
}

func BookRouter(route fiber.Router) {
	route.Get("/books", service.GetAllBooks)
	route.Get("/books/:id", service.GetBook)
	route.Put("/books/:id", service.UpdateBook)
	route.Post("/book", service.CreateBook)
	route.Delete("/books/:id", service.DeleteBook)
}
