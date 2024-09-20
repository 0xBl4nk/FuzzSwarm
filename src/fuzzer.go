package src

import (
    "log"
    "net/http"
    "strings"
    "sync"
    "time"
    "io/ioutil"
    "github.com/fatih/color"
)

func StartFuzzing(cfg Config) {
    log.Println("Starting fuzzing...")
    log.Printf("Using %d threads...", cfg.Threads)
    log.Printf("Total values to fuzz: %d", len(cfg.Values))

    var wg sync.WaitGroup
    semaphore := make(chan struct{}, cfg.Threads)

    client := CreateClient(cfg.UseProxy)

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

func FuzzRequest(cfg Config, client *http.Client, value string) {
    requestURL := strings.Replace(cfg.URL, "BRUTE", value, -1)
    req, err := http.NewRequest("GET", requestURL, nil)
    if err != nil {
        log.Printf("Failed to create request: %v", err)
        return
    }

    for headerKey, headerValue := range cfg.Headers {
        req.Header.Set(headerKey, headerValue)
    }

    if cfg.RateLimit > 0 {
        time.Sleep(time.Millisecond * time.Duration(cfg.RateLimit))
    }

    resp, err := client.Do(req)
    if err != nil {
        log.Printf("Request failed: %v", err)
        return
    }
    defer resp.Body.Close()

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Failed to read response body: %v", err)
        return
    }

    responseBody := string(bodyBytes)
    responseSize := len(responseBody)

    if cfg.FilterSize > 0 && responseSize == cfg.FilterSize {
        return
    }

    printResponse(cfg, value, resp.StatusCode, responseSize, responseBody)
}

func printResponse(cfg Config, value string, statusCode int, responseSize int, responseBody string) {
    if cfg.Verbose {
        previewLength := 100
        if len(responseBody) > previewLength {
            responseBody = responseBody[:previewLength] + "..."
        }

        if statusCode >= 200 && statusCode < 300 {
            color.New(color.FgGreen).Printf("Value: %s [%d] - Response Size: %d - Preview: %s\n", value, statusCode, responseSize, responseBody)
        } else if statusCode >= 300 && statusCode < 400 {
            color.New(color.FgYellow).Printf("Value: %s [%d] - Response Size: %d - Preview: %s\n", value, statusCode, responseSize, responseBody)
        } else {
            color.New(color.FgRed).Printf("Value: %s [%d] - Response Size: %d - Preview: %s\n", value, statusCode, responseSize, responseBody)
        }
    } else {
        if statusCode >= 200 && statusCode < 300 {
            color.New(color.FgGreen).Printf("Value: %s [%d] - Response Size: %d\n", value, statusCode, responseSize)
        } else if statusCode >= 300 && statusCode < 400 {
            color.New(color.FgYellow).Printf("Value: %s [%d] - Response Size: %d\n", value, statusCode, responseSize)
        } else {
            color.New(color.FgRed).Printf("Value: %s [%d] - Response Size: %d\n", value, statusCode, responseSize)
        }
    }
}

