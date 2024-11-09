package config

import (
    "os"
    "testing"
)

func TestLoadConfig(t *testing.T) {
    os.Setenv("DATABASE_URI", "mongodb://localhost:3000")
    os.Setenv("PORT", "8080")

    config := LoadConfig()

    if config.DatabaseURI != "mongodb://localhost:3000" {
        t.Errorf("Expected DatabaseURI to be 'mongodb://localhost:3000', got %s", config.DatabaseURI)
    }
    if config.Port != "8080" {
        t.Errorf("Expected Port to be '8080', got %s", config.Port)
    }
}
