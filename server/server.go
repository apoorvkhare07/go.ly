package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	app := fiber.New()
	app.Get("hello", helloWorldApi)
	return app
}

func Listen(app *fiber.App, port string) {
	err := app.Listen(port)
	if err != nil {
		fmt.Println("Error listening to port:", port)
		fmt.Println(err)
	}

}
