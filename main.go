package main

import (
	"github.com/Theshawa/syntax-highlighter/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	app := server.NewServer()
	err = app.Run()
	if err != nil {
		panic(err)
	}
}
