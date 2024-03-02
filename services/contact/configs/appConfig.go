package configs

import (
	"helloWRLDs/clean_arch/pkg/drivers/postgres"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Addr string
	Db   *postgres.PostgresCFG
}

func LoadConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		return nil
	}
	return &AppConfig{
		Addr: os.Getenv("ADDRESS"),
		Db: &postgres.PostgresCFG{
			Host:     os.Getenv("DBHOST"),
			Port:     os.Getenv("DBPORT"),
			User:     os.Getenv("DBUSER"),
			Password: os.Getenv("DBPASS"),
			Name:     os.Getenv("DBNAME"),
		},
	}
}
