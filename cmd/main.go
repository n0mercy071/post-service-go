package main

import (
	"flag"
	"log"

	postservicego "github.com/n0mercy071/post-service-go"
	handler "github.com/n0mercy071/post-service-go/pkg/handler"
	"github.com/n0mercy071/post-service-go/pkg/repository"
	"github.com/n0mercy071/post-service-go/pkg/service"
)

type ServerConfig struct {
	port string
}

func main() {
	var serverConfig ServerConfig
	flag.StringVar(&serverConfig.port, "port", "8000", "server port")
	flag.Parse()

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(postservicego.Server)

	if err := srv.Run(serverConfig.port, handlers.InitRoutes()); err != nil {
		log.Fatalln(err.Error())
	}
}
