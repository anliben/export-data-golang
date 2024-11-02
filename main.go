package main

import (
	"fmt"
	"log"
	"time"

	"github.com/anliben/export-data/pkg/excel"
	"github.com/anliben/export-data/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	excel.RegisterRoutes(app)

	go func() {
		for {
			fmt.Println("Limpando arquivos antigos...")
			err := utils.CleanOldFiles("./public")
			if err != nil {
				log.Printf("Erro na limpeza de arquivos: %v", err)
			}
			time.Sleep(24 * time.Hour)
		}
	}()

	app.Static("/", "./public")

	utils.StartServerWithGracefulShutdown(app)
}
