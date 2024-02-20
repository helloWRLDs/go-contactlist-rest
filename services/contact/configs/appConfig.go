package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Addr string
	Db   *DatabaseConfig
}

type DatabaseConfig struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
}

func LoadConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		return nil
	}
	return &AppConfig{
		Addr: os.Getenv("ADDRESS"),
		Db: &DatabaseConfig{
			Host:     os.Getenv("DBHOST"),
			Port:     os.Getenv("DBPORT"),
			Username: os.Getenv("DBUSER"),
			Password: os.Getenv("DBPASS"),
			Name:     os.Getenv("DBNAME"),
		},
	}
}
