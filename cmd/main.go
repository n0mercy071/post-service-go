package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	postservicego "github.com/n0mercy071/post-service-go"
	handler "github.com/n0mercy071/post-service-go/pkg/handler"
	"github.com/n0mercy071/post-service-go/pkg/repository"
	"github.com/n0mercy071/post-service-go/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title post-service-go
// @version 1.0

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey jwtToken
// @in header
// @name Authorization
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

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalln(err.Error())
		}
	}()

	logrus.Print("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App shutdown")

	if err := db.Close(); err != nil {
		logrus.Errorf("error db close: %s", err.Error())
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
