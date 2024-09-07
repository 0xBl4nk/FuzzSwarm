# FuzzSwarm

**FuzzSwarm** is a fuzzing tool designed for brute-forcing HTTP endpoints. It supports optional proxy usage, SSL configuration, and response size filtering to focus on significant results.

## 📃 Requirements

- `Python 3.10.x` or higher
- Python packages listed in `requirements.txt`

## Installation

1. Clone this repository:
    ```bash
    git clone https://github.com/your-username/FuzzSwarm.git
    ```

2. Install the dependencies:
    ```bash
    pip3 install -r requirements.txt
    ```

## Usage

To run FuzzSwarm, use the following syntax:

```bash
./FuzzSwarm <url> -H <headers.txt> --range <start-end> --ssl <ssl.pem> [options]
```

### Example Usage:

```bash
./FuzzSwarm -H headers.txt --range 001-1000 --ssl charles.pem --use-proxy https://example.com/api/BRUTE
```

### Available Parameters:

- `-H`: Path to the headers file.
- `--range`: Defines the numeric range for fuzzing (e.g., 001-100).
- `--wordlist`: Path to a wordlist file.
- `--ssl`: Path to the SSL certificate file.
- `--use-proxy`: Enables proxy configuration, which is loaded from the `.env` file.
- `-fs`: Filters out HTTP responses of a specific size. (skip responses with this size.)

## Contributing

1. Fork this repository.
2. Create a new branch: `git checkout -b <branch_name>`.
3. Make your changes and commit: `git commit -m '<commit_message>'`.
4. Push to your branch: `git push origin <branch_name>`.
5. Open a pull request.
