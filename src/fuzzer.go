package src

import (
    "io"
    "net/http"
    "strings"
    "sync"
    "time"

    "github.com/fatih/color"
)

// StartFuzzing initiates the fuzzing process based on the provided configuration.
func StartFuzzing(cfg Config) {
    LogInfo("Starting fuzzing...")
    LogInfo("Using %d threads...", cfg.Threads)
    LogInfo("Total values to fuzz: %d", len(cfg.Values))

    var wg sync.WaitGroup
    semaphore := make(chan struct{}, cfg.Threads)

    client := CreateClient(cfg.UseProxy, cfg.Timeout, cfg.SSLCertPath)

    for _, value := range cfg.Values {
        semaphore <- struct{}{}
        wg.Add(1)
        go func(val string) {
            defer func() {
                wg.Done()
                <-semaphore
            }()
            FuzzRequest(cfg, client, val)
        }(value)
    }

    wg.Wait()
}

// FuzzRequest performs the HTTP request for a single fuzzing value with retry logic.
func FuzzRequest(cfg Config, client *http.Client, value string) {
    requestURL := strings.Replace(cfg.URL, "BRUTE", value, -1)
    var req *http.Request
    var err error

    if cfg.Method == "POST" {
        fuzzedData := strings.ReplaceAll(cfg.Data, "BRUTE", value)
        req, err = http.NewRequest("POST", requestURL, strings.NewReader(fuzzedData))
        if err != nil {
            LogError("Failed to create POST request for value '%s': %v", value, err)
            return
        }
        req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    } else { // GET
        req, err = http.NewRequest("GET", requestURL, nil)
        if err != nil {
            LogError("Failed to create GET request for value '%s': %v", value, err)
            return
        }
    }

    for headerKey, headerValue := range cfg.Headers {
        req.Header.Set(headerKey, headerValue)
    }

    var resp *http.Response
    for attempt := 1; attempt <= cfg.Retries; attempt++ {
        if cfg.RateLimit > 0 {
            time.Sleep(time.Millisecond * time.Duration(cfg.RateLimit))
        }

        resp, err = client.Do(req)
        if err == nil {
            break
        }
        LogError("Request failed for value '%s' on attempt %d: %v", value, attempt, err)
        time.Sleep(time.Second * time.Duration(attempt)) // Exponential backoff
    }

    if err != nil {
        LogError("All retry attempts failed for value '%s': %v", value, err)
        return
    }
    defer resp.Body.Close()

    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        LogError("Failed to read response body for value '%s': %v", value, err)
        return
    }

    responseBody := string(bodyBytes)
    responseSize := len(responseBody)

    if cfg.FilterSize > 0 && responseSize == cfg.FilterSize {
        return
    }

    printResponse(cfg, value, resp.StatusCode, responseSize, responseBody)
}

// printResponse displays the HTTP response based on the configuration.
func printResponse(cfg Config, value string, statusCode int, responseSize int, responseBody string) {
    colorFunc := getColorFunc(statusCode)

    if cfg.Verbose {
        previewLength := 100
        if len(responseBody) > previewLength {
            responseBody = responseBody[:previewLength] + "..."
        }
        colorFunc.Printf("Value: %s [%d] - Response Size: %d - Preview: %s\n", value, statusCode, responseSize, responseBody)
    } else {
        colorFunc.Printf("Value: %s [%d] - Response Size: %d\n", value, statusCode, responseSize)
    }
}

// getColorFunc returns the appropriate color based on the HTTP status code.
func getColorFunc(statusCode int) *color.Color {
    switch {
    case statusCode >= 200 && statusCode < 300:
        return color.New(color.FgGreen)
    case statusCode >= 300 && statusCode < 400:
        return color.New(color.FgYellow)
    default:
        return color.New(color.FgRed)
    }
}
