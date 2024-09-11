# fuzzswarm

`fuzzswarm` is a customizable and powerful URL fuzzing tool designed to test API endpoints by manipulating injectable parameters using a wordlist or a range of values.

## Features

- **Support for Wordlist and Range:** You can provide a wordlist or a range of values for fuzzing.
- **Multithreading:** Control the number of threads to send requests in parallel.
- **Response Size Filtering:** Filter responses based on the size of the response body.
- **Rate Limiting:** Set a delay between requests to avoid overloading targets.
- **Color-Coded Status Codes:** Responses with different status codes are displayed in different colors for easy reading.

## Requirements

- Go 1.16 or higher

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/fuzzswarm.git
   cd fuzzswarm
  ```
Build the tool:

bash

    go build

Usage

You can use fuzzswarm to fuzz URLs with a wordlist or a range of values.
Example of Fuzzing with a Range:

bash

./fuzzswarm --range 1-10000,3 -t 10 -url http://192.168.1.35:5000/api/BRUTE -rl 100

In this example, the tool:

    Uses a range of values from 1 to 10000, with 3 digits (e.g., 001, 002).
    Utilizes 10 threads to send requests in parallel.
    The target endpoint is http://192.168.1.35:5000/api/BRUTE, where BRUTE will be replaced by the fuzzing values.
    Sets a rate limit of 100 milliseconds between each request.

Example of Fuzzing with a Wordlist:

bash

./fuzzswarm -w wordlist.txt -t 10 -url http://192.168.1.35:5000/api/BRUTE -rl 100

In this case, the tool:

    Reads fuzzing values from wordlist.txt.
    Uses 10 threads and a rate limit of 100 milliseconds between requests.

Options:

    -url: Target URL, where BRUTE is the fuzzing injection point.
    -t: Number of threads.
    -rl: Rate limit between requests in milliseconds.
    -fs: Response size filter (ignores responses of the specified size).
    -w: Path to the wordlist.
    --range: Number range to be used for fuzzing (e.g., 1-1000,3).

Contributing

    Fork this repository.
    Create a new feature branch: git checkout -b my-new-feature
    Commit your changes: git commit -m 'Add new feature'
    Push to the branch: git push origin my-new-feature
    Open a pull request!

License

This project is licensed under the MIT License - see the LICENSE file for details.

markdown


### Key Sections:

- **Overview:** Explains the purpose of the tool.
- **Features:** Lists the main features of the tool.
- **Installation:** Explains how to clone and build the project.
- **Usage:** Shows examples of how to run the tool with a wordlist or range.
- **Options:** Describes the command-line options available.
- **Contributing:** Explains how to contribute to the project.
- **License:** Mentions the licensing information.

Now your project is fully documented in English! If you need further adjustments
