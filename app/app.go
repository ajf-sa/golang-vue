package main

import (
	"os"

	"github.com/alfuhigi/store-ajf-sa/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(postgres.Open(helpers.Config("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if e := server(db); e != nil {
		panic(e)
	}
}

func server(db *gorm.DB) error {
	config := fiber.Config{
		ServerHeader: "go",
	}

	for i := range os.Args[1:] {
		if os.Args[1:][i] == "-prefork" {
			config.Prefork = true
		}
	}

	app := fiber.New(config)
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	app.Static("/static", "./static/dist")

	app.Static("/", "./webapp")

	app.Get("/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./webapp/index.html")
	})

	return app.Listen(":3000")

}
