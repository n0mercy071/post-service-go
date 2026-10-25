package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	postservicego "github.com/n0mercy071/post-service-go"
	handler "github.com/n0mercy071/post-service-go/pkg/handler"
	"github.com/n0mercy071/post-service-go/pkg/repository"
	"github.com/n0mercy071/post-service-go/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error config: %s", err.Error())
	}

	db, err := initDB()
	if err != nil {
		log.Fatalf("error init db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(postservicego.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalln(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func initDB() (*sqlx.DB, error) {
	return repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("postgres_host"),
		Port:     viper.GetString("postgres_port"),
		Username: viper.GetString("postgres_user"),
		Password: viper.GetString("postgres_password"),
		DBName:   viper.GetString("postgres_db"),
		SSLMode:  viper.GetString("postgres_sslmode"),
	})
}
