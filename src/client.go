package src

import (
    "crypto/tls"
    "net/http"
    "net/url"
    "os"
    "time"
)

// CreateClient configures and returns an HTTP client.
// It sets up proxy and SSL certificates if provided.
func CreateClient(useProxy bool, timeout int, sslCertPath string) *http.Client {
    transport := &http.Transport{}

    if useProxy {
        proxy := os.Getenv("HTTP_PROXY")
        if proxy == "" {
            LogFatal("Proxy configuration is missing in the .env file.")
        }

        proxyURL, err := url.Parse(proxy)
        if err != nil {
            LogFatal("Invalid proxy URL: %v", err)
        }
        transport.Proxy = http.ProxyURL(proxyURL)
        LogInfo("Using proxy: %s", proxy)

        // SSL certificate setup if sslCertPath is provided
        if sslCertPath != "" {
            cert, err := tls.LoadX509KeyPair(sslCertPath, sslCertPath)
            if err != nil {
                LogFatal("Failed to load SSL certificate from %s: %v", sslCertPath, err)
            }
            transport.TLSClientConfig = &tls.Config{Certificates: []tls.Certificate{cert}}
            LogInfo("Using SSL certificate from: %s", sslCertPath)
        }
    }

    client := &http.Client{
        Transport: transport,
        Timeout:   time.Duration(timeout) * time.Second,
    }

    return client
}
