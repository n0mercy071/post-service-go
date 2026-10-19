package main

import (
	"flag"
	"log"

	postservicego "github.com/n0mercy071/post-service-go"
	handler "github.com/n0mercy071/post-service-go/pkg/handler"
)

type ServerConfig struct {
	port string
}

func main() {
	var serverConfig ServerConfig
	flag.StringVar(&serverConfig.port, "port", "8000", "server port")
	flag.Parse()

	handlers := new(handler.Handler)
	srv := new(postservicego.Server)

	if err := srv.Run(serverConfig.port, handlers.InitRoutes()); err != nil {
		log.Fatalln(err.Error())
	}
}
