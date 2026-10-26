package main

import (
	"github.com/jmoiron/sqlx"
	postservicego "github.com/n0mercy071/post-service-go"
	handler "github.com/n0mercy071/post-service-go/pkg/handler"
	"github.com/n0mercy071/post-service-go/pkg/repository"
	"github.com/n0mercy071/post-service-go/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error config: %s", err.Error())
	}

	db, err := initDB()
	if err != nil {
		logrus.Fatalf("error init db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(postservicego.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalln(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func initDB() (*sqlx.DB, error) {
	return repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
}
