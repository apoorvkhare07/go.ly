package main

import (
	"fmt"
	"github.com/apoorvkhare07/go.ly/model"
	"github.com/apoorvkhare07/go.ly/server"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	err := godotenv.Load(".env")
	environment := os.Getenv("Environment")
	fmt.Println(environment)
	//err = sentry.Init(sentry.ClientOptions{
	//	Dsn:              os.Getenv("SentryDsn"),
	//	Environment:      environment,
	//	TracesSampleRate: 1.0,
	//})

	if err != nil {
		panic(err)
	}

	//defer sentry.Flush(2 * time.Second)
	//defer sentry.Recover()

	config := &model.Config{
		Host:     os.Getenv("DBHost"),
		Port:     os.Getenv("DBPort"),
		Password: os.Getenv("DBPassword"),
		User:     os.Getenv("DBUser"),
		SSLMode:  "disable",
		DBName:   os.Getenv("DBName"),
	}

	err = model.Setup(config)
	if err != nil {
		panic(err)
	}
	app := server.Setup()

	server.Listen(app)

	//if environment == "production" {
	//	Init(app)
	//	lambda.Start(Handler)
	//} else {
	//	server.Listen(app)
	//}
}
