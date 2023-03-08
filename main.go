package main

import (
	"fmt"
	"github.com/apoorvkhare07/go.ly/server"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	fmt.Println("Hello world!")
	_ = godotenv.Load(".env")
	port := os.Getenv("port")
	app := server.Setup()
	server.Listen(app, ":"+port)
}
