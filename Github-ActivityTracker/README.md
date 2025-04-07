# GitHub Activity Tracker

GitHub Activity Tracker is a Go application that allows you to quickly check a GitHub user's recent activity directly from your terminal. It uses GitHub's public API to fetch user events and presents them in a readable format.

## Features

- Fetch the most recent GitHub activities of any public user
- Display activities with timestamps
- Support for various GitHub event types (PushEvent, CreateEvent, IssueEvent, etc.)
- Simple interface with clear error handling
- No external dependencies - uses only Go standard library

## Installation

### Prerequisites

- Go 1.16 or higher

### Building from Source

1. Clone this repository or download the source code:

```bash
git clone https://github.com/yourusername/github-activity-cli.git
cd github-activity-cli
```

2. Build the application:

```bash
go build -o github-activity main.go
```

3. (Optional) Move the binary to a directory in your PATH for easier access:

```bash
sudo mv github-activity /usr/local/bin/
```

## Usage

Run the application by providing a GitHub username as an argument:

```bash
github-activity aakaru
```

Example output:

```
Recent GitHub Activity:
----------------------
[Apr 01, 2025 14:32]  aakaru pushed to octocat/Hello-World
[Mar 30, 2025 09:15]  aakaru worked on a pull request in octocat/Hello-World
[Mar 29, 2025 17:45]  aakaru commented on an issue in octocat/Hello-World
[Mar 28, 2025 11:23]  aakaru updated an issue in org/repo-name
[Mar 27, 2025 16:08]  aakaru starred popular/repository
```

## Error Handling

The CLI handles common errors such as:
- Invalid usernames
- Network connection issues
- API request failures

## License

[MIT License](LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
