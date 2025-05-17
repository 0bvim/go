# FT_MONITORING

A simple yet powerful website monitoring tool written in Go that checks the status of websites and logs their availability.

## Overview

This project is a command-line application that allows you to:
- Monitor the status of multiple websites defined in a `sites.txt` file
- Check if websites are online (HTTP 200 response)
- Save logs of website availability with timestamps
- View monitoring history through a simple interface

## Features

- **Website Status Checking**: Performs HTTP requests to verify website availability
- **Configurable Monitoring**: Checks each site multiple times with configurable delay between checks
- **Logging System**: Records monitoring results with timestamps in a structured format
- **Simple CLI Interface**: Easy-to-use menu-based interface with clear options
- **Persistent History**: Maintains a log file that persists between program runs

## Installation

### Clone the repository:

```bash
git clone git@github.com:vinicius-f-pereira/go.git
```

Or using GitHub CLI:

```bash
gh repo clone vinicius-f-pereira/go
```

### Navigate to the project directory:

```bash
cd go/basic/ft_monitoring
```

## Usage

### Run the application:

```bash
go run monitoring.go
```

### Build an executable:

```bash
go build monitoring.go
```

Then run the compiled binary:

```bash
./monitoring
```

### Program Flow:

When you run the application, you'll see a welcome message with the version number, followed by the menu options.

## Menu Options

When you run the application, you'll be presented with the following options:

1. **Check Status** - Checks the status of all websites in `sites.txt`
   - Makes HTTP requests to each website
   - Displays the status in real-time
   - Records results to the log file

2. **Show Logs** - Displays the monitoring history from the log file
   - Shows timestamp, website URL, and online status for each check

0. **Exit** - Terminates the application

## Configuration

### Website List

Add the websites you want to monitor to the `sites.txt` file, with one URL per line:

```
https://example.com
https://another-site.com
https://your-site.org
```

The application comes pre-configured with several example websites.

### Monitoring Parameters

You can modify the following constants in the code:
- `check` (default: 2) - Number of monitoring cycles
- `delay` (default: 5) - Delay between checks in seconds

To change these values, modify the constants at the top of `monitoring.go`:

```go
const (
    check = 2    // Number of monitoring cycles
    delay = 5    // Delay between checks in seconds
)
```

## How It Works

1. The application reads URLs from `sites.txt`
2. For each URL, it sends an HTTP GET request
3. It logs the response status (200 = online, otherwise = not working)
4. Results are stored in `log.txt` with timestamps
5. The monitoring cycle repeats based on the `check` constant
6. Between each website check, the program waits for the duration specified by `delay`

## Log Format

Logs are stored in `log.txt` with the following format:
```
MM/DD/YYYY HH:MM:SSAM/PM - https://example.com - online: true/false
```

Example log entry:
```
10/11/2023 01:31:38AM - https://intra.42.fr/ - online: true
```

## Project Structure

- `monitoring.go` - Main application code
- `sites.txt` - List of websites to monitor
- `log.txt` - Log file with monitoring history
- `monitoring_commented.go` - Commented version with explanations (for learning purposes)

## Development Notes

This is a learning project built to explore Go programming concepts including:
- HTTP requests and responses
- File I/O operations (reading, writing, appending)
- Error handling and logging
- Type definitions and constants
- String formatting and manipulation
- Time management and formatting
- Command-line user interface

## Future Enhancements

Potential improvements that could be added:
- Configuration file for monitoring parameters
- Support for HTTP methods beyond GET
- Response time tracking
- Alert system for when websites go down
- Web interface for viewing monitoring data
- Support for authentication for protected sites

## License

This project is open source and available under the [MIT License](LICENSE).