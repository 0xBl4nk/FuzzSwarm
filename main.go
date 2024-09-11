package main

import (
    "fuzzswarm/src"
    "log"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables from the .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found or failed to load. Continuing without environment variables.")
    }

    config, err := src.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    src.StartFuzzing(config)
}

