# OBSOLETE version! Access FuzzSwarm2 in my repositories, or through the link:  https://github.com/0xbl4nk/FuzzSwarm2
OBSOLETE version! Access FuzzSwarm2 in my repositories, or through the link:  https://github.com/0xbl4nk/FuzzSwarm2

# FuzzSwarm

**FuzzSwarm** is a fuzzing tool designed for brute-forcing HTTP endpoints. It supports optional proxy usage, SSL configuration, and response size filtering to focus on significant results.

## 📃 Requirements

- `Go 1.23.1` or higher

## Installation

1. Clone this repository:
    ```bash
    git clone https://github.com/your-username/FuzzSwarm.git
    ```

2. Install the dependencies:
    ```bash
    make build
    ```

## Usage

To run FuzzSwarm, use the following syntax:

```bash
./FuzzSwarm -url <url> -range/-w
```

### Example Usage:

```bash
./fuzzswarm -R 1-10000,3 -t 10 -X POST -d "param=BRUTE"-u http://192.168.1.35:5000/api/test -f 34 -v
```
<img src="https://i.imgur.com/8sf7iLI.png">

### Available Parameters:

```
Flags:
  -d, --data string       POST data with 'BRUTE' as the placeholder for injection.
  -f, --filter-size int   Filter responses by size (skip responses with this size).
  -H, --headers string    Optional path to the headers file.
  -h, --help              help for fuzzswarm
  -X, --method string     HTTP method to use (GET or POST). (default "GET")
  -R, --range string      Range of numbers to use, format start-end,digits (e.g., 1-10000,3).
  -r, --rate-limit int    Rate limit in milliseconds between requests.
      --ssl-cert string   Path to SSL certificate file (optional).
  -t, --threads int       Number of threads to use for fuzzing. (default 10)
  -u, --url string        The target URL with 'BRUTE' as the placeholder for injection points.
  -p, --use-proxy         Enable proxy configuration from .env file.
  -v, --verbose           Display verbose output including response preview.
  -w, --wordlist string   Path to the wordlist file.
```

# How to Generate a Valid SSL Certificate with OpenSSL
fuzzswarm uses unified certificates, i.e. key and certificate in the same file.

```bash
cat certificate.pem privatekey.pem > fullcert.pem
```
If your unified certificate is .c12, you will need to convert to pem.
```bash
openssl pkcs12 -in yourfile.p12 -clcerts -nokeys -out certificate.pem
openssl pkcs12 -in yourfile.p12 -nocerts -out privatekey.pem
openssl rsa -in privatekey.pem -out privatekey.pem
cat certificate.pem privatekey.pem > fullcert.pem
```

## Contributing

1. Fork this repository.
2. Create a new branch: `git checkout -b <branch_name>`.
3. Make your changes and commit: `git commit -m '<commit_message>'`.
4. Push to your branch: `git push origin <branch_name>`.
5. Open a pull request.

