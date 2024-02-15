package main

import (
	pkg "helloWRLDs/clean_arch/pkg/store/postgres"
	"helloWRLDs/clean_arch/services/contact/internal/delivery"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"
)

type application struct {
	delivery delivery.HTTPDeliveryInterface
	logger   *slog.Logger
}

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	var (
		dbHost     string = viper.GetString("database.host")
		dbPort     string = viper.GetString("database.port")
		dbUser     string = viper.GetString("database.user")
		dbPassword string = viper.GetString("database.password")
		dbName     string = viper.GetString("database.name")
		port       string = viper.GetString("server.address")
	)
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	db, err := pkg.GetPostgresConnection(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	} else {
		logger.Info("connected to db", "host", dbHost, "port", dbPort, "db", dbName)
	}

	app := application{
		delivery: delivery.NewDelivery(db, logger),
		logger:   logger,
	}

	srv := &http.Server{
		Handler:      app.routes(),
		Addr:         port,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.logger.Info("server started", "port", port)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
