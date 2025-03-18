# Reminder Tool

A command-line reminder tool written in Go that allows you to set timed notifications with custom messages.

## Overview

`remind-me` is a lightweight CLI application that lets you schedule reminder notifications. Simply specify a time and your reminder message, and the application will display a system notification when the time arrives.

## Features

- Simple command-line interface
- Natural time input (e.g., "14:30", "2:45pm")
- System notifications using [beeep](https://github.com/gen2brain/beeep)
- Cross-platform support (Windows, macOS, Linux)
- Non-blocking operation (continues running in background)

## Installation

### Prerequisites

- Go 1.16 or higher
- Git

### Install from source

```bash
# Clone the repository
git clone https://github.com/yourusername/remind-me.git
cd remind-me

# Install dependencies
go get -u github.com/gen2brain/beeep
go get -u github.com/olebedev/when

# Build the application
go build -o remind-me

# Optional: Move to a directory in your PATH
mv remind-me /usr/local/bin/  # Linux/macOS
# or
# move remind-me C:\Windows\System32\  # Windows
```

## Usage

```bash
remind-me <time> <message>
```

### Examples

```bash
# Set a reminder for 3:30 PM
remind-me 15:30 Take the cookies out of the oven

# Set a reminder for 5 minutes from now
remind-me "5 minutes from now" Call back John

# Set a reminder for tomorrow morning
remind-me "tomorrow at 9am" Team meeting
```

## How It Works

The application uses the [when](https://github.com/olebedev/when) library to parse natural language time expressions. When executed, it calculates the time difference between now and the specified time, then spawns a background process that will display a notification using the [beeep](https://github.com/gen2brain/beeep) library when that time arrives.

## Dependencies

- [github.com/gen2brain/beeep](https://github.com/gen2brain/beeep) - Cross-platform notification library
- [github.com/olebedev/when](https://github.com/olebedev/when) - Natural language date/time parser

## License

[MIT License](LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4
