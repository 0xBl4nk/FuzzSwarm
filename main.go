package main

import (
    "FuzzSwarm/src"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/spf13/cobra"
)

func main() {
    var rootCmd = &cobra.Command{
        Use:   "fuzzswarm",
        Short: "FuzzSwarm is a fuzzing tool for testing URLs",
        Run: func(cmd *cobra.Command, args []string) {
            // Load environment variables from the .env file
            err := godotenv.Load()
            if err != nil {
                log.Println("No .env file found or failed to load. Continuing without environment variables.")
            }

            config, err := src.LoadConfig(cmd)
            if err != nil {
                src.LogFatal("Failed to load configuration: %v", err)
            }

            src.StartFuzzing(config)
        },
    }

    // Define flags directly on the root command
    rootCmd.Flags().StringP("url", "u", "", "The target URL with 'BRUTE' as the placeholder for injection points.")
    rootCmd.Flags().StringP("headers", "H", "", "Optional path to the headers file.")
    rootCmd.Flags().StringP("wordlist", "w", "", "Path to the wordlist file.")
    rootCmd.Flags().BoolP("use-proxy", "p", false, "Enable proxy configuration from .env file.")
    rootCmd.Flags().IntP("threads", "t", 10, "Number of threads to use for fuzzing.")
    rootCmd.Flags().IntP("filter-size", "f", 0, "Filter responses by size (skip responses with this size).")
    rootCmd.Flags().IntP("rate-limit", "r", 0, "Rate limit in milliseconds between requests.")
    rootCmd.Flags().StringP("range", "R", "", "Range of numbers to use, format start-end,digits (e.g., 1-10000,3).")
    rootCmd.Flags().BoolP("verbose", "v", false, "Display verbose output including response preview.")
    rootCmd.Flags().StringP("method", "X", "GET", "HTTP method to use (GET or POST).")
    rootCmd.Flags().StringP("data", "d", "", "POST data with 'BRUTE' as the placeholder for injection.")

    // SSL certificate flag
    rootCmd.Flags().String("ssl-cert", "", "Path to SSL certificate file (optional).")

    if err := rootCmd.Execute(); err != nil {
        log.Println(err)
        os.Exit(1)
    }
}
