package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host            string
	Port            string
	Password        string
	User            string
	DBName          string
	SSLMode         string
	MigrationTarget string
}

func GetConfig() Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	config := Config{
		Host:            os.Getenv("DB_HOST"),
		Port:            os.Getenv("DB_PORT"),
		Password:        os.Getenv("DB_PASS"),
		User:            os.Getenv("DB_USER"),
		SSLMode:         os.Getenv("DB_SSLMODE"),
		DBName:          os.Getenv("DB_NAME"),
		MigrationTarget: os.Getenv("MIGRATION_TARGET"),
	}

	return config
}
