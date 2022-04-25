package main

import (
	"gitlab.com/Valghall/diwor/server/cmd"
	"gitlab.com/Valghall/diwor/server/internal/handler"
	"gitlab.com/Valghall/diwor/server/internal/service"
	storage2 "gitlab.com/Valghall/diwor/server/internal/storage"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfigs(); err != nil {
		logrus.Fatalf("Error while parsing configs: %v", err)
	}

	err := godotenv.Load()
	if err != nil {
		logrus.Errorf("Error while parsing .env: %v", err)
	}

	var db *sqlx.DB
	if os.Getenv("ENV") == "dev" {
		db, err = storage2.NewPostgresDB(storage2.Config{
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
		db, err = storage2.NewPostgresDBFromURL()
		if err != nil {
			logrus.Fatalln(err)
		}
	}

	store := storage2.NewStorage(db)
	services := service.NewServices(store)
	handlers := handler.NewHandler(services)
	server := new(cmd.Server)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = viper.GetString("port")
	}
	log.Fatalln(server.Run(port, handlers.InitRoutes()))
}

func initConfigs() error {
	viper.AddConfigPath("server/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
