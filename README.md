# Advent of Code

Solutions for Advent of Code challenges implemented primarily in Go.

## Prerequisites

- Go 1.21.4 or higher
- A valid Advent of Code session cookie

## Installation

1. Clone the repository

2. Set up your Advent of Code session cookie as an environment variable:

```bash
export AOC_SESSION=your_session_cookie_here
```

You can find your session cookie by:

1. Logging into [Advent of Code](https://adventofcode.com)
2. Opening browser developer tools
3. Going to Application/Storage > Cookies
4. Copying the value of the 'session' cookie

## Running Solutions

Run a specific day's solution:

```bash
go run cmd/main.go <day>
```

For example, to run day 1:

```bash
go run cmd/main.go 1
```

If no day is specified, it will default to day 1.

## Project Structure

- `/cmd/main.go` - Main entry point
- `/internal/dayN` - Solutions for each day
- `/inputs` - Input files (automatically downloaded and cached)

## Testing

Run tests for a specific day:

```bash
go test ./internal/day1
```

Run all tests:

```bash
go test ./...
```

## Features

- Automatic input fetching and caching
- Performance timing for solutions
- Modular structure for easy addition of new days
