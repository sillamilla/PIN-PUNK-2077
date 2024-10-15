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
	HTTP        HTTP
	HackService HackService
}

type HTTP struct {
	Port string
}

type HackService struct {
	HackServiceAddress string
	Endpoint           string
}

func GetConfig() *Config {
	if c == nil {
		if err := godotenv.Load("../../.env"); err != nil {
			log.Println("Error loading .env file")
		}

		// HTTP
		httpPort := os.Getenv("PORT")
		if httpPort == "" {
			httpPort = "8080"
			log.Println("PORT is not set, using default: 8080")
		}

		// HACK SERVICE
		hackServiceAddress := os.Getenv("HACK_SERVICE_PORT")
		if hackServiceAddress == "" {
			hackServiceAddress = "http://localhost:8081"
			log.Println("HACK_SERVICE_PORT is not set, using default: http://localhost:8081")
		}

		endpoint := os.Getenv("ENDPOINT")
		if endpoint == "" {
			endpoint = "/hack"
			log.Println("PORT is not set, using default: /hack")
		}

		c = &Config{
			HTTP: HTTP{
				Port: httpPort,
			},
			HackService: HackService{
				HackServiceAddress: hackServiceAddress,
				Endpoint:           endpoint,
			},
		}
	}

	return c
}
