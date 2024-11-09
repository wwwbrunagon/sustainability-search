package config

import (
    "os"
    "log"
    "github.com/joho/godotenv"
)

type Config struct {
    DatabaseURI string
    Port        string
}

func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using system environment variables")
    }

    return &Config{
        DatabaseURI: os.Getenv("DATABASE_URI"),
        Port:        os.Getenv("PORT"),
    }
}
