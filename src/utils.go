package src

import (
    "bufio"
    "os"
    "strings"
)

func ReadHeaders(path string) (map[string]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    headers := make(map[string]string)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.SplitN(line, ": ", 2)
        if len(parts) == 2 {
            headers[parts[0]] = parts[1]
        }
    }

    return headers, scanner.Err()
}

func ReadValues(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var values []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        values = append(values, scanner.Text())
    }

    return values, scanner.Err()
}

