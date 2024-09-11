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
    go build
    ```

## Usage

To run FuzzSwarm, use the following syntax:

```bash
./FuzzSwarm -url <url> -range/-w
```

### Example Full Usage:

```bash
./fuzzswarm -range 1-10000,3 -t 10 -url http://192.168.1.35:5000/api/BRUTE -rl 6 -use-proxy -fs 34
```
<img src="https://i.imgur.com/m1wXsMB.png">

### Available Parameters:

- `-H`: Path to the headers file.
- `-range`: Range of numbers to use, format start-end,digits (e.g., 1-10000,3).
- `-w`: Path to a wordlist file.
- `-use-proxy`: Enable proxy and SSL configuration from .env file.
- `-fs`: Filters out HTTP responses of a specific size. (skip responses with this size.)
- `-t`: Number of threads to use for fuzzing.

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
