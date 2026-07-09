package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key     string
	BaseUrl string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("no dotenv var")
	}

	key := os.Getenv("KEY")
	baseUrl := os.Getenv("BASE_URL")

	if key == "" || baseUrl == "" {
		panic("variable env required")
	}
	return &Config{
		Key:     key,
		BaseUrl: baseUrl,
	}
}
