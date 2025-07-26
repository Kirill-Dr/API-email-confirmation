package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Verify  VerifyConfig
	Project ProjectConfig
}

type ProjectConfig struct {
	PORT string
}

type VerifyConfig struct {
	Email    string
	Password string
	Address  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	return &Config{
		Project: ProjectConfig{
			PORT: os.Getenv("PORT"),
		},
		Verify: VerifyConfig{
			Email:    os.Getenv("EMAIL"),
			Password: os.Getenv("PASSWORD"),
			Address:  os.Getenv("ADDRESS"),
		},
	}
}
