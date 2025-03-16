# Hello, Earth!

This CLI tool is a simple demonstration of how to use Go's core packages: testing, fmt, and flag.

## Features

- **fmt:** Demonstrates formatted I/O.
- **flag:** Parses command-line options.
- **testing:** Provides examples of how to write and run tests with Go.

## Installation

Clone the repository and build the project:

```bash
git clone https://github.com/yourusername/hello-earth.git
cd hello-earth
go build -o hello-earth
```

## Usage

Run the tool to see available commands and options:

```bash
./hello-earth --help
```

## Running Tests

Execute the tests using Go's test command:

```bash
go test ./...
```

## Structure

- `main.go`: Entry point of the CLI application.
- `tests/`: Contains test files demonstrating Go's testing package.
- Additional files showcase examples of using the fmt and flag packages.

## License

Distributed under the MIT License.