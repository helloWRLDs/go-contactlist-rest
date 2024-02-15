package main

import (
	"fmt"
	pkg "helloWRLDs/clean_arch/pkg/store/postgres"
	"helloWRLDs/clean_arch/services/contact/internal/delivery"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"
)

type application struct {
	delivery delivery.HTTPDeliveryInterface
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
	db, err := pkg.GetPostgresConnection(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := application{
		delivery: delivery.NewDelivery(db),
	}

	srv := &http.Server{
		Handler:      app.routes(),
		Addr:         port,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("server started on port ", port)
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
