package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	c *Config
)

type Config struct {
	Postgres Postgres
	HTTP     HTTP
}

type HTTP struct {
	Port string
}

type Postgres struct {
	Addr     string
	User     string
	Password string
	DBName   string
}

func GetConfig() *Config {
	if c == nil {
		if err := godotenv.Load("../../.env"); err != nil {
			log.Println("Error loading .env file")
		}

		// Postgres
		addr := os.Getenv("POSTGRES_ADDR")
		if addr == "" {
			addr = "localhost:5432"
			log.Println("POSTGRES_ADDR is not set, using default: localhost:5432")

		}

		user := os.Getenv("POSTGRES_USER")
		if user == "" {
			user = "postgres"
			log.Println("POSTGRES_USER is not set, using default: postgres")
		}

		password := os.Getenv("POSTGRES_PASSWORD")
		if password == "" {
			password = "postgres"
			log.Println("POSTGRES_PASSWORD is not set, using default: postgres")
		}

		dbName := os.Getenv("POSTGRES_DB")
		if dbName == "" {
			dbName = "postgres"
			log.Println("POSTGRES_DB is not set, using default: postgres")
		}

		// HTTP
		httpPort := os.Getenv("PORT")
		if httpPort == "" {
			httpPort = "8081"
			log.Println("PORT is not set, using default: 8081")
		}

		c = &Config{
			Postgres: Postgres{
				Addr:     addr,
				User:     user,
				Password: password,
				DBName:   dbName,
			},
			HTTP: HTTP{
				Port: httpPort,
			},
		}
	}

	return c
}
