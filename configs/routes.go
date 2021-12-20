package configs

import (
	"github.com/nucarlos/go-bender/controllers"

	swagger "github.com/arsmn/fiber-swagger/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	// Add info endpoint to routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API Falabella FIF - 🐣 v1.0.0")
	})

	app.Get("/version", func(c *fiber.Ctx) error {
		return c.SendString("API Falabella FIF - 🐣 v1.0.0")
	})

	// Add swagger handler to routes as a middleware
	app.Get("/docs/*", swagger.Handler)
	app.Get("/docs/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
	}))

	// Add api endpoint to routes
	api := app.Group("/api", logger.New())
	beers := api.Group("/beers")
	beers.Get("/", controllers.GetAllBeers)
	beers.Post("/", controllers.CreateBeers)
	beers.Get("/:beerID", controllers.GetBeer)
	beers.Get("/:beerID/boxprice", controllers.GetBeerBoxPrice)

	// Register route for 404 Error.
	notFoundRoute(app)

}

// notFoundRoute func for describe 404 Error route.
func notFoundRoute(a *fiber.App) {
	a.Use(
		// Anonimus function.
		func(c *fiber.Ctx) error {
			// Return HTTP 404 status and JSON response.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "sorry, endpoint is not found",
			})
		},
	)
}
