package main

import (
	"log"
	"os"

	"gitlab.com/Valghall/diwor/cmd"
	"gitlab.com/Valghall/diwor/internal/handler"
)

func main() {
	handlers := new(handler.Handler)
	server := new(cmd.Server)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	log.Fatalln(server.Run(port, handlers.InitRoutes()))
}
