# FT_MONITORING

A simple website monitoring tool written in Go that checks the status of websites and logs their availability.

## Overview

This project is a command-line application that allows you to:
- Monitor the status of multiple websites defined in a `sites.txt` file
- Check if websites are online (HTTP 200 response)
- Save logs of website availability
- View monitoring history

## Features

- **Website Status Checking**: Performs HTTP requests to verify website availability
- **Configurable Monitoring**: Checks each site multiple times with configurable delay
- **Logging System**: Records monitoring results with timestamps
- **Simple CLI Interface**: Easy-to-use menu-based interface

## Installation

### Clone the repository:

```link
git clone git@github.com:vinicius-f-pereira/go.git
```

Or using GitHub CLI:

```link
gh repo clone vinicius-f-pereira/go
```

### Navigate to the project directory:

```
cd go/src/ft_monitoring
```

## Usage

### Run the application:

```
go run monitoring.go
```

### Build an executable:

```
go build monitoring.go
```

Then run the compiled binary:

```
./monitoring
```

## Menu Options

When you run the application, you'll be presented with the following options:

1. **Check Status** - Checks the status of all websites in `sites.txt`
2. **Show Logs** - Displays the monitoring history from the log file
0. **Exit** - Terminates the application

## Configuration

### Website List

Add the websites you want to monitor to the `sites.txt` file, with one URL per line:

```
https://example.com
https://another-site.com
```

### Monitoring Parameters

You can modify the following constants in the code:
- `check` (default: 2) - Number of monitoring cycles
- `delay` (default: 5) - Delay between checks in seconds

## How It Works

1. The application reads URLs from `sites.txt`
2. For each URL, it sends an HTTP GET request
3. It logs the response status (200 = online, otherwise = not working)
4. Results are stored in `log.txt` with timestamps

## Log Format

Logs are stored in the following format:
```
MM/DD/YYYY HH:MM:SSAM/PM - https://example.com - online: true/false
```

## Development Notes

This is a learning project built to explore Go programming concepts including:
- HTTP requests
- File I/O operations
- Error handling
- Type definitions
- String formatting
- Time management

## License

This project is open source and available under the [MIT License](LICENSE).