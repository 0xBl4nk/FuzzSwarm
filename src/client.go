package src

import (
    "crypto/tls"
    "log"
    "net/http"
    "net/url"
    "os"
)

func CreateClient(useProxy bool) *http.Client {
    client := &http.Client{}

    if useProxy {
        proxy := os.Getenv("HTTP_PROXY")
        var proxyURL *url.URL
        var err error

        if proxy != "" {
            proxyURL, err = url.Parse(proxy)
            if err != nil {
                log.Fatalf("Invalid proxy URL: %v", err)
            }
            client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
            log.Printf("Using proxy: %s", proxy)
        } else {
            log.Println("Proxy configuration is missing in the .env file.")
            log.Println("Exiting program since -use-proxy flag was used but no proxy is configured.")
            os.Exit(1)
        }

        sslCertPath := os.Getenv("SSL_CERT_PATH")
        if sslCertPath != "" {
            cert, err := tls.LoadX509KeyPair(sslCertPath, sslCertPath)
            if err != nil {
                log.Fatalf("Failed to load SSL certificate from %s: %v", sslCertPath, err)
            }
            client.Transport = &http.Transport{
                Proxy:           http.ProxyURL(proxyURL),
                TLSClientConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
            }
            log.Printf("Using SSL certificate from: %s", sslCertPath)
        } else {
            log.Println("SSL certificate configuration is missing in the .env file.")
            log.Println("Exiting program since -use-proxy flag was used but no SSL certificate is configured.")
            os.Exit(1)
        }
    }

    return client
}

