package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/nucarlos/go-bender/configs"
	"github.com/nucarlos/go-bender/utils"
	"github.com/urfave/cli"

	_ "github.com/nucarlos/go-bender/docs"
)

var (
	app *cli.App
)

func init() {
	// dotenv
	err := godotenv.Load()
	if err != nil {
		log.Print(err.Error())
	}
}

// @title API Falabella FIF
// @version 1.0.0
// @description Esta API esta diseñada para ser una prueba para los nuevos candidatos al equipo.
// @contact.name nucarlos
// @contact.email nucarlos@gmail.com
// @BasePath /
func main() {
	// Initialise a CLI app
	app = cli.NewApp()
	app.Name = "API Falabella FIF"
	app.Usage = "Challenge Test Falabella"
	app.Description = "Esta API esta diseñada para ser una prueba para los nuevos candidatos al equipo."
	app.Author = "nucarlos"
	app.Email = "nucarlos@gmail.com"
	app.Version = "1.0.0"
	app.Compiled = time.Now()
	app.Copyright = "(c) 2021 nucarlos"
	app.Action = func(c *cli.Context) error {
		serverRun()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func serverRun() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Fiber instance
	app := fiber.New(config)
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/Caracas",
	}))

	// DB instance
	configs.Database()

	// Setting routes
	configs.SetupRoutes(app)

	// Start server
	utils.StartServer(app)
}
