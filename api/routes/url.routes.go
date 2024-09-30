package routes

import (
	"url-shortner/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UrlRoutes(r *fiber.App) {
	route := r.Group("/api/v1/url")

	//for shortening urls
	route.Post("/short", controllers.ShortenUrl)

	//for redirecting using short urls
	route.Get("/:id", controllers.RedirectToUrl)
}
