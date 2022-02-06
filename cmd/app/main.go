package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	"gitlab.com/Valghall/diwor/internal/service"

	_ "github.com/lib/pq"
	"gitlab.com/Valghall/diwor/internal/storage"

	"gitlab.com/Valghall/diwor/cmd"
	"gitlab.com/Valghall/diwor/internal/handler"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfigs(); err != nil {
		logrus.Fatalf("Error while parsing configs: %v", err)
	}

	err := godotenv.Load()
	if err != nil {
		logrus.Errorf("Error while parsing .env: %v", err)
	}

	var db *sqlx.DB
	if os.Getenv("ENV") == "dev" {
		db, err = storage.NewPostgresDB(storage.Config{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
			Password: os.Getenv("DB_PASSWORD"),
		})
		if err != nil {
			logrus.Fatalln(err)
		}
	} else {
		db, err = storage.NewPostgresDBFromURL()
		if err != nil {
			logrus.Fatalln(err)
		}
	}

	store := storage.NewStorage(db)
	service := service.NewService(store)
	handlers := handler.NewHandler(service)
	server := new(cmd.Server)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = viper.GetString("port")
	}
	log.Fatalln(server.Run(port, handlers.InitRoutes()))
}

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
