package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func (c *Config) NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("no dotenv var")
	}
	key := os.Getenv("KEY")
	fmt.Println(key)
	if key == "" {
		panic("Key variable env required")
	}
	return &Config{
		Key: key,
	}
}
