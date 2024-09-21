package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

type Config struct {
	ProjectID string
	GeminiAPIKey       string
	NewsAPIKey string
	ServiceAccountKeyPath string
	DB_NAME string
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using system environment variables")
    }

    config := &Config{
		ProjectID: os.Getenv("PROJECT_ID"),
		GeminiAPIKey: os.Getenv("GEMINI_API_KEY"),
		NewsAPIKey: os.Getenv("NEWS_API_KEY"),
		ServiceAccountKeyPath: os.Getenv("SERVICE_ACCOUNT_KEY_PATH"),
		DB_NAME: os.Getenv("DB_NAME"),
    }

    return config, nil
}
