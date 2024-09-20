package src

import (
    "flag"
    "os"
    "strings"
    "strconv"
    "fmt"
)

type Config struct {
    URL        string
    Headers    map[string]string
    UseProxy   bool
    Values     []string
    Threads    int
    FilterSize int
    RateLimit  int
    Range      string
    Verbose    bool
}

func LoadConfig() (Config, error) {
    var cfg Config
    flag.StringVar(&cfg.URL, "url", "", "The target URL with 'BRUTE' as the placeholder for injection points.")
    headersPath := flag.String("H", "", "Optional path to the headers file.")
    wordlistPath := flag.String("w", "", "Path to the wordlist file.")
    flag.BoolVar(&cfg.UseProxy, "use-proxy", false, "Enable proxy and SSL configuration from .env file.")
    flag.IntVar(&cfg.Threads, "t", 10, "Number of threads to use for fuzzing.")
    flag.IntVar(&cfg.FilterSize, "fs", 0, "Filter responses by size (skip responses with this size).")
    flag.IntVar(&cfg.RateLimit, "rl", 0, "Rate limit in milliseconds between requests.")
    flag.StringVar(&cfg.Range, "range", "", "Range of numbers to use, format start-end,digits (e.g., 1-10000,3).")
    flag.BoolVar(&cfg.Verbose, "v", false, "Display verbose output including response preview.")
    flag.Parse()

    if cfg.URL == "" {
        flag.Usage()
        os.Exit(1)
    }

    if *headersPath != "" {
        headers, err := ReadHeaders(*headersPath)
        if err != nil {
            return cfg, err
        }
        cfg.Headers = headers
    }

    if *wordlistPath != "" {
        values, err := ReadValues(*wordlistPath)
        if err != nil {
            return cfg, err
        }
        cfg.Values = values
    } else if cfg.Range != "" {
        values, err := parseRange(cfg.Range)
        if err != nil {
            return cfg, err
        }
        cfg.Values = values
    } else {
        return cfg, fmt.Errorf("Either a range or a wordlist must be provided")
    }

    return cfg, nil
}

// Função para parsear o range
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

