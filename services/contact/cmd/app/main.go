package main

import (
	"fmt"
	"helloWRLDs/clean_arch/pkg/drivers/postgres"
	"helloWRLDs/clean_arch/services/contact/configs"
	"helloWRLDs/clean_arch/services/contact/internal/delivery/httpDelivery"
	"net/http"
	"os"
	"time"
)

func main() {
	cfg := configs.LoadConfig()

	db, err := postgres.GetPostgresConnection(cfg.Db)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Connected to db")

	httpRouter := httpDelivery.NewHttpRouter(db)
	httpRouter.InitRoutes()

	srv := &http.Server{
		Handler:      httpRouter.Router,
		Addr:         cfg.Addr,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server started")
	if err := srv.ListenAndServe(); err != nil {
		os.Exit(1)
	}

}
