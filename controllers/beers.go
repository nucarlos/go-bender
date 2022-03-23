package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nucarlos/go-bender/models"
	"github.com/nucarlos/go-bender/services"
)

// GetAllBeers
// @Summary Lista todas las cervezas
// @Description Lista todas las cervezas que se encuentran en la base de datos
// @Tags cerveza
// @Accept json
// @Produce json
// @Success 200 {object} []models.BeerItem "Operacion exitosa"
// @Router /api/beers [get]
func GetAllBeers(ctx *fiber.Ctx) error {
	// Beer Service instance
	beerService := services.NewBeerService()

	beers, err := beerService.SearchBeers()
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"beer_items": *beers,
	})
}

// CreateBeers
// @Summary Ingresa una nueva cerveza
// @Description Ingresa una nueva cerveza
// @Tags cerveza
// @Accept json
// @Produce json
// @Param id body integer true "ID de la cerveza"
// @Param name body string true "Nombre de la cerveza"
// @Param brewery body string true "Cervecería"
// @Param country body string true "Pais de origen"
// @Param price body number true "Precio"
// @Param currency body string true "Tipo de moneda"
// @Success 200 {object} models.BeerItem "Operacion exitosa"
// @Failure 400 {string} string "Request invalida"
// @Failure 409 {string} string "El ID de la cerveza ya existe"
// @Router /api/beers [post]
func CreateBeers(ctx *fiber.Ctx) error {
	// Beer Service instance
	beerService := services.NewBeerService()

	// Validate params
	request := new(models.BeerItemPayload)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).SendString("Request invalida 1")
	}
	if err := request.Validate(); err != nil {
		return ctx.Status(400).SendString("Request invalida 2")
	}

	if err := beerService.AddBeers(request.ID, request.Name, request.Brewery, request.Country, request.Currency, request.Price); err != nil {
		return ctx.SendStatus(500)
	}

	return ctx.SendStatus(200)
}

// GetBeer
// @Summary Lista el detalle de la marca de cervezas
// @Description Busca una cerveza por su Id
// @Tags cerveza
// @Accept json
// @Produce json
// @Param beerID path integer true "Busca una cerveza por su Id"
// @Success 200 {object} models.BeerItem "Operacion exitosa"
// @Failure 404 {string} string "El Id de la cerveza no existe"
// @Router /api/beers/{beerID} [get]
func GetBeer(ctx *fiber.Ctx) error {
	// Beer Service instance
	beerService := services.NewBeerService()

	var id int64
	if _, err := fmt.Sscan(ctx.Params("beerID"), &id); err != nil {
		return ctx.Status(400).SendString("Request invalida 1")
	}

	beer, err := beerService.SearchBeerById(id)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"beer_item": beer,
	})
}

// GetBeerBoxPrice
// @Summary Lista el precio de una caja de cervezas de una marca
// @Description Obtiene el precio de una caja de cerveza por su Id
// @Tags cerveza
// @Accept json
// @Produce json
// @Param beerID path integer true "Busca una cerveza por su Id"
// @Param currency query string true "Tipo de moneda con la que pagará"
// @Param quantity query integer false "La cantidad de cervezas a comprar"
// @Success 200 {object} models.BeerBox "Operacion exitosa"
// @Failure 404 {string} string "El Id de la cerveza no existe"
// @Router /api/beers/{beerID}/boxprice [get]
func GetBeerBoxPrice(ctx *fiber.Ctx) error {
	// Beer Service instance
	beerService := services.NewBeerService()

	// Validate params
	var id int64
	if _, err := fmt.Sscan(ctx.Params("beerID"), &id); err != nil {
		return ctx.Status(400).SendString("Request invalida")
	}
	payload := new(models.BeerBoxPayload)
	if err := ctx.QueryParser(payload); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	beerBox, err := beerService.BoxBeerPriceById(id, payload.Currency, payload.Quantity)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"beer_box": beerBox,
	})
}
