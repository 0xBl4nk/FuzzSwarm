package src

import (
    "errors"
    "fmt"
    "net/url"
    "strconv"
    "strings"

    "github.com/spf13/cobra"
)

// Config holds all configuration settings for the fuzzing process.
type Config struct {
    URL         string
    Headers     map[string]string
    UseProxy    bool
    Values      []string
    Threads     int
    FilterSize  int
    RateLimit   int
    Range       string
    Verbose     bool
    Timeout     int
    Retries     int
    Method      string
    Data        string
    SSLCertPath string
}

// LoadConfig parses and validates command-line flags to populate the Config struct.
func LoadConfig(cmd *cobra.Command) (Config, error) {
    var cfg Config

    cfg.URL, _ = cmd.Flags().GetString("url")
    headersPath, _ := cmd.Flags().GetString("headers")
    wordlistPath, _ := cmd.Flags().GetString("wordlist")
    cfg.UseProxy, _ = cmd.Flags().GetBool("use-proxy")
    cfg.Threads, _ = cmd.Flags().GetInt("threads")
    cfg.FilterSize, _ = cmd.Flags().GetInt("filter-size")
    cfg.RateLimit, _ = cmd.Flags().GetInt("rate-limit")
    cfg.Range, _ = cmd.Flags().GetString("range")
    cfg.Verbose, _ = cmd.Flags().GetBool("verbose")
    cfg.Timeout = 10 // Default timeout in seconds
    cfg.Retries = 3  // Default number of retries
    cfg.Method, _ = cmd.Flags().GetString("method")
    cfg.Data, _ = cmd.Flags().GetString("data")
    cfg.SSLCertPath, _ = cmd.Flags().GetString("ssl-cert")

    if cfg.URL == "" {
        return cfg, errors.New("the --url flag is required")
    }

    // Validate URL format
    parsedURL, err := url.Parse(cfg.URL)
    if err != nil || !(parsedURL.Scheme == "http" || parsedURL.Scheme == "https") {
        return cfg, errors.New("invalid URL format. Ensure it starts with http:// or https://")
    }

    // Validate HTTP method
    cfg.Method = strings.ToUpper(cfg.Method)
    if cfg.Method != "GET" && cfg.Method != "POST" {
        return cfg, errors.New("invalid HTTP method. Only GET and POST are supported")
    }

    // If method is POST, data must be provided
    if cfg.Method == "POST" && cfg.Data == "" {
        return cfg, errors.New("POST method requires --data flag")
    }

    if headersPath != "" {
        headers, err := ReadHeaders(headersPath)
        if err != nil {
            return cfg, fmt.Errorf("error reading headers: %v", err)
        }
        cfg.Headers = headers
    }

    if wordlistPath != "" {
        values, err := ReadValues(wordlistPath)
        if err != nil {
            return cfg, fmt.Errorf("error reading wordlist: %v", err)
        }
        cfg.Values = values
    } else if cfg.Range != "" {
        values, err := parseRange(cfg.Range)
        if err != nil {
            return cfg, fmt.Errorf("error parsing range: %v", err)
        }
        cfg.Values = values
    } else {
        return cfg, errors.New("either a range or a wordlist must be provided")
    }

    return cfg, nil
}

// parseRange generates a slice of strings based on the provided range string.
func parseRange(rangeStr string) ([]string, error) {
    parts := strings.Split(rangeStr, ",")
    if len(parts) != 2 {
        return nil, fmt.Errorf("range format should be start-end,digits (e.g., 1-10000,3)")
    }
    rangeParts := strings.Split(parts[0], "-")
    if len(rangeParts) != 2 {
        return nil, fmt.Errorf("range bounds format should be start-end")
    }
    start, err := strconv.Atoi(rangeParts[0])
    if err != nil {
        return nil, fmt.Errorf("invalid start value: %v", err)
    }
    end, err := strconv.Atoi(rangeParts[1])
    if err != nil {
        return nil, fmt.Errorf("invalid end value: %v", err)
    }
    digits, err := strconv.Atoi(parts[1])
    if err != nil {
        return nil, fmt.Errorf("invalid digits value: %v", err)
    }

    var values []string
    for i := start; i <= end; i++ {
        values = append(values, fmt.Sprintf("%0*d", digits, i))
    }
    return values, nil
}
