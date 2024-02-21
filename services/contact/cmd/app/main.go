package main

import (
	"fmt"
	pkg "helloWRLDs/clean_arch/pkg/store/postgres"
	"helloWRLDs/clean_arch/services/contact/configs"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type application struct {
	logger *slog.Logger
}

func main() {
	cfg := configs.LoadConfig()
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	db, err := pkg.GetPostgresConnection(
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.Username,
		cfg.Db.Password,
		cfg.Db.Name,
	)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	} else {
		logger.Info("connected to db", "url", fmt.Sprintf("'%s:%s/%s'", cfg.Db.Host, cfg.Db.Port, cfg.Db.Name))
	}

	app := application{
		logger: logger,
	}

	srv := &http.Server{
		Handler:      app.routes(),
		Addr:         cfg.Addr,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.logger.Info("server started", "port", cfg.Addr)
	if err = srv.ListenAndServe(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
