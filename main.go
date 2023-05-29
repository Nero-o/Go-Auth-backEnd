package main

import (
	"crud_go/database"
	"crud_go/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))

	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("CORS Config:", c.Response().Header.Peek("Access-Control-Allow-Origin"))
		return c.Next()
	})

	routes.Setup(app)

	err = app.Listen(":8000")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
